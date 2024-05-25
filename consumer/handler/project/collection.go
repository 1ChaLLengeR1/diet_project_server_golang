package project

import (
	"database/sql"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionProject struct{
	Collection []project_data.Collection 	`json:"collection"`
	Status     int 								`json:"status"` 
	Pagination helpers.PaginationCollectionPost `json:"pagination"`
	Error      string 							`json:"error"`
}


func HandlerCollectionProject(c *gin.Context) {
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Query: c.Query("private"),
		Param: c.Param("page"),
	}


	collection, err := CollectionProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCollectionProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Pagination: collection.Pagination,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCollectionProject{
		Collection: collection.Collection,
		Pagination: collection.Pagination,
		Status: collection.Status,
		Error: collection.Error,
	})
}

func CollectionProject(params params_data.Params)(ResponseCollectionProject, error) {
	userData := params.Header
    queryParam := params.Query


    var usersData []user_data.User
    var collectionsData []project_data.Collection
    var query string

    perPage := 16
    page := 1

    db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionProject{}, err
    }

	if queryParam == "true" {
        _, users, err := auth.CheckUser(userData)
        if err != nil {
            return ResponseCollectionProject{}, err
        }
        usersData = users

        query = `SELECT * FROM project WHERE "userId" = $1 ORDER BY "createdUp" DESC LIMIT $2 OFFSET $3;`
    } else {
        query = `SELECT * FROM project ORDER BY "createdUp" DESC LIMIT $1 OFFSET $2;`
    }

	pageStr := params.Param
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "project", page, perPage)

	var rows *sql.Rows
    if queryParam == "true" {
        rows, err = db.Query(query, &usersData[0].Id, perPage, pagination.Offset)
    } else {
        rows, err = db.Query(query, perPage, pagination.Offset)
    }

	if err != nil {
		return ResponseCollectionProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection project_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.Title, &collection.Description, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseCollectionProject{}, err
		}
		collectionsData = append(collectionsData, collection)
	}

	return ResponseCollectionProject{
		Collection: collectionsData,
		Status: http.StatusOK,
		Pagination: pagination,
		Error: "",
	}, nil
}