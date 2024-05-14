package post

import (
	"database/sql"
	params_data "internal/consumer/data"
	collection_data "internal/consumer/data/post"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"internal/consumer/handler/auth"
	helpers "internal/consumer/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseCollection struct{
	Collection []collection_data.Collection
	Status     int
	Pagination helpers.PaginationCollectionPost
	Error      string
}

func HandlerCollection(c *gin.Context){

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Query: c.Query("private"),
		Param: c.Param("page"),
	}


	collection, err := collection(c, params)
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

func collection(c* gin.Context, params params_data.Params)(ResponseCollection, error){
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
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp, &collection.Description); err != nil {
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