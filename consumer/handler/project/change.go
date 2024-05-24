package project

import (
	"fmt"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseChnageProject struct {
	Collection []project_data.Change 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}


func HandlerChangeProject(c *gin.Context) {
	var changePost project_data.Change
	c.BindJSON(&changePost)
	jsonMap, err := helpers.BindJSONToMap(&changePost)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseChnageProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
		
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("changeId"),
		Json: jsonMap,
	}

	change, err := ChangeProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseChnageProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseChnageProject{
		Collection: change.Collection,
		Status: change.Status,
		Error: change.Error,
	})
}

func ChangeProject(params params_data.Params)(ResponseChnageProject, error){
	userData := params.Header
	
	var usersData []user_data.User
	var changesData []project_data.Change


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseChnageProject{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseChnageProject{}, err
	}

	usersData = users
	id := params.Param

	title, titleOk := params.Json["title"].(string)
	description, descriptionOk := params.Json["description"].(string)
	updateUp, updateUpOk := params.Json["updateUp"].(string)
	
	var updateFields []string
	if titleOk {
		updateFields = append(updateFields, fmt.Sprintf(`"title"='%s'`, title)) 
	}
	if descriptionOk {
		updateFields = append(updateFields, fmt.Sprintf(`"description"='%s'`, description))
	}
	if updateUpOk {
		updateFields = append(updateFields, fmt.Sprintf(`"updateUp"='%s'`, updateUp)) 
	}

	if len(updateFields) == 0 {
		if err != nil {
			return ResponseChnageProject{}, err
		}
	}

	query := `UPDATE project SET` +  strings.Join(updateFields, ", ") + ` WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", "title", "description", "createdUp", "updateUp";`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseChnageProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var change project_data.Change
		if err := rows.Scan(&change.Id, &change.UserId, &change.Title, &change.Description, &change.CreatedUp, &change.UpdateUp); err != nil {
			return ResponseChnageProject{}, err
		}
		changesData = append(changesData, change)
	}

	return ResponseChnageProject{
		Collection: changesData,
		Status: http.StatusOK,
		Error: "",
	}, nil

}
