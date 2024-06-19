package post

import (
	"database/sql"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	training_data "myInternal/consumer/data/training"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	training_function "myInternal/consumer/handler/training"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionOne struct{
	Collection []post_data.Collection `json:"collection"`
	CollectionTraining []training_data.Collection `json:"collectionTraining"`
	Status     int							`json:"status"`
	Error      string 						`json:"error"`
}

func HandlerCollectionOne(c *gin.Context){

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Query: c.Query("private"),
		Param: c.Param("id"),
	}

	collectionOne, err := CollectionOne(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollectionOne{
			Collection: nil,
			CollectionTraining: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	params = params_data.Params{
		Param: collectionOne.Collection[0].Id,
	}

	collectionOneTraining, err := training_function.CollectionOneTraining(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollectionOne{
			Collection: nil,
			CollectionTraining: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}


	c.JSON(http.StatusOK, ResponseCollectionOne{
		Collection: collectionOne.Collection,
		CollectionTraining: collectionOneTraining.Collection,
		Status: collectionOne.Status,
		Error: collectionOne.Error,
	})
}

func CollectionOne(params params_data.Params)(ResponseCollectionOne, error){
	userData := params.Header
	queryParam := params.Query

	var usersData []user_data.User
	var collectionOneData []post_data.Collection
	var query string


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOne{}, err
	}
	defer db.Close()

	if queryParam == "true" {
        _, users, err := auth.CheckUser(userData)
        if err != nil {
            return ResponseCollectionOne{}, err
        }
        usersData = users

        query = `SELECT * FROM post WHERE "id" = $1 AND "userId" = $2;`
    } else {
        query = `SELECT * FROM post WHERE "id" = $1`
    }

	id := params.Param


	var rows *sql.Rows
    if queryParam == "true" {
        rows, err = db.Query(query, &id, &usersData[0].Id)
    } else {
        rows, err = db.Query(query, &id)
    }

	if err != nil {
		return ResponseCollectionOne{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection post_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionOne{}, err
		}
		collectionOneData = append(collectionOneData, collection)
	}

	return ResponseCollectionOne{
		Collection: collectionOneData,
		CollectionTraining: nil,
		Status: http.StatusOK,
		Error: "",
	}, nil
}