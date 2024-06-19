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
		AppLanguage: c.GetHeader("AppLanguage"),
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
	appLanguage := params.AppLanguage

    var usersData []user_data.User
    var collectionsData []project_data.Collection
    var query string

    perPage := 16
    page := 1

    db, err := database.ConnectToDataBase()
    if err != nil {
        return ResponseCollectionProject{}, err
    }
	defer db.Close()

	if queryParam == "true" {
        _, users, err := auth.CheckUser(userData)
        if err != nil {
            return ResponseCollectionProject{}, err
        }
        usersData = users

        query = `WITH filtered_projects AS (
			SELECT * 
			FROM project 
			WHERE "userId" = $1 
			ORDER BY "createdUp" DESC 
			LIMIT $2 OFFSET $3
		)
		SELECT 
			p.id, 
			p."userId", 
			pml."idLanguage", 
			pml.title, 
			pml.description, 
			p."createdUp", 
			p."updateUp"
		FROM 
			filtered_projects p
		JOIN 
			project_multi_language pml ON p.id = pml."idProject"
		WHERE 
			pml."idLanguage" = $4;
		`
    } else {
        query = `WITH limited_projects AS (
			SELECT * 
			FROM project 
			ORDER BY "createdUp" DESC 
			LIMIT $1 OFFSET $2
		)
		SELECT 
			lp.id, 
			lp."userId", 
			pml."idLanguage", 
			pml.title, 
			pml.description, 
			lp."createdUp", 
			lp."updateUp"
		FROM 
			limited_projects lp
		JOIN 
			project_multi_language pml ON lp.id = pml."idProject"
		WHERE 
			pml."idLanguage" = $3;`
    }

	pageStr := params.Param
    if pageStr != "" {
        page, _ = strconv.Atoi(pageStr)
    }

	pagination := helpers.GetPaginationData(db, "project", page, perPage)

	var rows *sql.Rows
    if queryParam == "true" {
        rows, err = db.Query(query, &usersData[0].Id, perPage, pagination.Offset, appLanguage)
    } else {
        rows, err = db.Query(query, perPage, pagination.Offset, appLanguage)
    }
	defer rows.Close()

	if err != nil {
		return ResponseCollectionProject{}, err
	}
	
	for rows.Next() {
		var collection project_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.IdLanguage, &collection.Title, &collection.Description, &collection.CreatedUp, &collection.UpdateUp); err != nil {
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