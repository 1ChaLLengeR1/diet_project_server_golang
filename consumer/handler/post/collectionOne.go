package post

import (
	collection_data "internal/consumer/data/post"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"internal/consumer/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCollectionOne struct{
	Collection []collection_data.Collection
	Status     int
	Error      string
}

func HandlerCollectionOne(c *gin.Context){
	collectionOne, err := collectionOne(c)
	if err != nil{
		c.JSON(http.StatusOK, ResponseCollectionOne{
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

func collectionOne(c *gin.Context)(ResponseCollectionOne, error){
	userData := c.GetHeader("UserData")
	var usersData []user_data.User
	var collectionOneData []collection_data.Collection


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCollectionOne{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCollectionOne{}, err
	}

	usersData = users
	id := c.Param("id")

	query := `SELECT * FROM post WHERE "id" = $1 AND "userId" = $2;`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseCollectionOne{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection collection_data.Collection
		if err := rows.Scan(&collection.Id, &collection.UserId, &collection.Day, &collection.Weight, &collection.Kcal, &collection.CreatedUp, &collection.UpdateUp, &collection.Description); err != nil {
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