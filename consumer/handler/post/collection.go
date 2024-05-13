package post

import (
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
	collection, err := collection(c)
	if err != nil{
		c.JSON(http.StatusOK, ResponseCollection{
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

func collection(c* gin.Context)(ResponseCollection, error){
	userData := c.GetHeader("UserData")
	var usersData []user_data.User
	var collectionsData []collection_data.Collection
	perPage := 16
	page := 1
	


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollection{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCollection{}, err
	}

	usersData = users
	
	pageStr := c.Param("page")
	if pageStr != ""{
		page, _ = strconv.Atoi(pageStr)
	}

	pagination := helpers.GetPaginationData(db, page, perPage)

	query := `SELECT * FROM post WHERE "userId" = $3 ORDER BY "updateUp" DESC LIMIT $1 OFFSET $2;`
	rows, err := db.Query(query, &perPage, pagination.Offset, &usersData[0].Id)
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