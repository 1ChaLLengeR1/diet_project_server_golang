package post

import (
	"database/sql"
	params_data "myInternal/consumer/data"
	collection_data "myInternal/consumer/data/post"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionOne struct{
	Collection []collection_data.Collection `json:"collection"`
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
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCollectionOne{
		Collection: collectionOne.Collection,
		Status: collectionOne.Status,
		Error: collectionOne.Error,
	})
}

func CollectionOne(params params_data.Params)(ResponseCollectionOne, error){
	userData := params.Header
	queryParam := params.Query

	var usersData []user_data.User
	var collectionOneData []collection_data.Collection
	var query string


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOne{}, err
	}

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
		var collection collection_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionOne{}, err
		}
		collectionOneData = append(collectionOneData, collection)
	}

	return ResponseCollectionOne{
		Collection: collectionOneData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}