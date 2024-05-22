package post

import (
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCreate struct {
	Collection []post_data.Post `json:"collection"`
	Status     int 				`json:"status"`
	Error      string 			`json:"error"`
}


func CreateHandler(c * gin.Context){

	var createPost post_data.Post
	c.BindJSON(&createPost)

	jsonMap, err := helpers.BindJSONToMap(&createPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseCreate{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
		
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
		Json: jsonMap,
	}

	craete, err := Create(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCreate{
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


func Create(params params_data.Params) (ResponseCreate, error){
	userData := params.Header
	var usersData []user_data.User
	var postsData []post_data.Post


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreate{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCreate{}, err
	}

	usersData = users

	day := params.Json["day"]
	projectId := params.Param
	weight := params.Json["weight"]
	kcal := params.Json["kcal"]
	createdUp := params.Json["createdUp"]
	updateUp := params.Json["updateUp"]
	description := params.Json["description"]


	query := `INSERT INTO post ("userId", "projectId", day, weight, kcal, "createdUp", "updateUp", description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING "id", "userId", "projectId", "day", "weight", "kcal", "createdUp", "updateUp", "description";`

	rows, err := db.Query(query, usersData[0].Id, projectId, day, weight, kcal, createdUp, updateUp, description)
	if err != nil {
		return ResponseCreate{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var post post_data.Post
		if err := rows.Scan(&post.Id, &post.UserId, &post.ProjectId, &post.Day, &post.Weight, &post.Kcal, &post.CreatedUp, &post.UpdateUp, &post.Description); err != nil {
			return ResponseCreate{}, err
		}
		postsData = append(postsData, post)
	}

	return ResponseCreate{
		Collection: postsData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}