package project

import (
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseCreateProject struct {
	Collection []project_data.Create 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}


func HandlerCreateProject(c *gin.Context){
	var createProject project_data.Create
	c.BindJSON(&createProject)

	jsonMap, err := helpers.BindJSONToMap(&createProject)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseCreateProject{
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

	project, err := CreateProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseCreateProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseCreateProject{
		Collection: project.Collection,
		Status: project.Status,
		Error: project.Error,
	})


}


func CreateProject(params params_data.Params)(ResponseCreateProject, error) {
	userData := params.Header
	var usersData []user_data.User
	var projectsData []project_data.Create

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseCreateProject{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseCreateProject{}, err
	}

	usersData = users

	title := params.Json["title"]
	description := params.Json["description"]
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")

	query := `INSERT INTO project ("userId", title, description, "createdUp", "updateUp") VALUES ($1, $2, $3, $4, $5) RETURNING "id", "userId", "title", "description", "createdUp", "updateUp";`

	rows, err := db.Query(query, usersData[0].Id, title, description, formattedDate, formattedDate)
    if err != nil {
        return ResponseCreateProject{}, err
    }
	defer rows.Close()

	for rows.Next() {
		var project project_data.Create
		if err := rows.Scan(&project.Id, &project.UserId, &project.Title, &project.Description, &project.CreatedUp, &project.UpdateUp); err != nil {
			return ResponseCreateProject{}, err
		}
		projectsData = append(projectsData, project)
	}

	return ResponseCreateProject{
		Collection: projectsData,
		Status: http.StatusOK,
		Error: "",
	}, nil

}