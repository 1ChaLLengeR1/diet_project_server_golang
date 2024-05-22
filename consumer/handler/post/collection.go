package post

import (
	"database/sql"
	params_data "myInternal/consumer/data"
	collection_data "myInternal/consumer/data/post"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseCollection struct{
	Collection []collection_data.Collection 	`json:"collection"`
	Status     int 								`json:"status"` 
	Pagination helpers.PaginationCollectionPost `json:"pagination"`
	Error      string 							`json:"error"`
}



func HandlerCollection(c *gin.Context){

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Query: c.Query("private"),
		Param: c.Param("page"),
	}


	collection, err := Collection(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollection{
			Collection: nil,
			Status: http.StatusBadRequest,
			Pagination: collection.Pagination,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCollection{
		Collection: collection.Collection,
		Pagination: collection.Pagination,
		Status: collection.Status,
		Error: collection.Error,
	})
}

func Collection(params params_data.Params)(ResponseCollection, error){
	userData := params.Header
    queryParam := params.Query


    var usersData []user_data.User
    var collectionsData []collection_data.Collection
    var query string

    perPage := 16
    page := 1

    db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollection{}, err
    }

	if queryParam == "true" {
        _, users, err := auth.CheckUser(userData)
        if err != nil {
            return ResponseCollection{}, err
        }
        usersData = users

        query = `SELECT * FROM post WHERE "userId" = $1 ORDER BY "day" DESC LIMIT $2 OFFSET $3;`
    } else {
        query = `SELECT * FROM post ORDER BY "day" DESC LIMIT $1 OFFSET $2;`
    }

	pageStr := params.Param
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "post", page, perPage)

	var rows *sql.Rows
    if queryParam == "true" {
        rows, err = db.Query(query, &usersData[0].Id, perPage, pagination.Offset)
    } else {
        rows, err = db.Query(query, perPage, pagination.Offset)
    }

	if err != nil {
		return ResponseCollection{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection collection_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.ProjectId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp, &collection.Description); err != nil {
			return ResponseCollection{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollection{
		Collection: collectionsData,
		Status: http.StatusOK,
		Pagination: pagination,
		Error: "",
	}, nil
}