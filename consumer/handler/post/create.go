package post

import (
	post_data "internal/consumer/data/post"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"internal/consumer/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCreate struct {
	Collection []post_data.Post
	Status     int
	Error      string
}


func CreateHandler(c * gin.Context){

	err, craete := create(c)
	if err != nil{
		c.JSON(http.StatusOK, ResponseCreate{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCreate{
		Collection: craete.Collection,
		Status: craete.Status,
		Error: craete.Error,
	})
}


func create(c *gin.Context) (error, ResponseCreate){
	userData := c.GetHeader("UserData")
	var usersData []user_data.User
	var postsData []post_data.Post


	db, err := database.ConnectToDataBase()
	if err != nil{
		return err, ResponseCreate{}
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return err, ResponseCreate{}
	}

	usersData = users


	var createPost post_data.Post
	err = c.BindJSON(&createPost)
	if err != nil{
		return err, ResponseCreate{}
	}



	query := `INSERT INTO post ("userId", day, weight, kcal, "createdUp", "updateUp", description) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING "id", "userId", "day", "weight", "kcal", "createdUp", "updateUp", "description";`

	rows, err := db.Query(query, usersData[0].Id, createPost.Day, createPost.Weight, createPost.Kcal, createPost.CreatedUp, createPost.UpdateUp, createPost.Description)
	if err != nil {
		return err, ResponseCreate{}
	}
	defer rows.Close()

	for rows.Next() {
		var post post_data.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.Day, &post.Weight, &post.Kcal, &post.CreatedUp, &post.UpdateUp, &post.Description); err != nil {
			return err, ResponseCreate{}
		}
		postsData = append(postsData, post)
	}

	return nil, ResponseCreate{
		Collection: postsData,
		Status: http.StatusOK,
		Error: "",
	}
}