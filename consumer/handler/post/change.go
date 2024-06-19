package post

import (
	"fmt"
	params_data "myInternal/consumer/data"
	change_data "myInternal/consumer/data/post"
	training_data "myInternal/consumer/data/training"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	training_function "myInternal/consumer/handler/training"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseChange struct{
	Collection []change_data.Change `json:"collection"`
	CollectionTraining []training_data.Change `json:"collectionTraining"`
	Status     int              	`json:"status"`
	Error      string          		`json:"error"`
}

func responseStatus(c *gin.Context, col []change_data.Change, colTrai []training_data.Change, status int, err error) {
	response := ResponseChange{
		Collection:         col,
		CollectionTraining: colTrai,
		Status:             status,
	}
	
	if err != nil {
		response.Error = err.Error()
	}
	
	c.JSON(status, response)
}

func HandlerChange(c *gin.Context){

	var changePost change_data.Change
	c.BindJSON(&changePost)
	jsonMap, err := helpers.BindJSONToMap(&changePost)
	if err != nil {
		responseStatus(c, nil, nil, http.StatusBadRequest, err)
		return
	}
		
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("id"),
		Json: jsonMap,
	}


	change, err := Change(params)
	if err != nil{
		responseStatus(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	params = params_data.Params{
		Param: change.Collection[0].Id,
		Json: jsonMap,
	}

	changeTraining, err := training_function.ChangeTraining(params)
	if err != nil{
		responseStatus(c, nil, nil, http.StatusBadRequest, err)
		return
	}

	responseStatus(c, change.Collection, changeTraining.Collection, change.Status, nil)
}

func Change(params params_data.Params)(ResponseChange, error){
	userData := params.Header
	
	var usersData []user_data.User
	var changesData []change_data.Change


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseChange{}, err
	}
	defer db.Close()

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseChange{}, err
	}

	usersData = users

	id := params.Param

	day, dayOk := params.Json["day"].(float64) 
	weight, weightOk := params.Json["weight"].(float64)
	kcal, kcalOk := params.Json["kcal"].(float64)
	now := time.Now()
    formattedDate := now.Format("2006-01-02 15:04:05")

	var updateFields []string
	if dayOk {
		updateFields = append(updateFields, fmt.Sprintf(`"day"=%d`, int64(day))) 
	}
	if weightOk {
		updateFields = append(updateFields, fmt.Sprintf(`"weight"=%f`, weight))
	}
	if kcalOk {
		updateFields = append(updateFields, fmt.Sprintf(`"kcal"=%d`, int64(kcal))) 
	}
	updateFields = append(updateFields, fmt.Sprintf(`"updateUp"='%s'`, formattedDate))
	
	if len(updateFields) == 0 {
		if err != nil {
			return ResponseChange{}, err
		}
	}

	query := `UPDATE post SET` +  strings.Join(updateFields, ", ") + ` WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", "projectId", "day", "weight", "kcal", "createdUp", "updateUp";`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseChange{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var change change_data.Change
		if err := rows.Scan(&change.Id, &change.UserId, &change.ProjectId, &change.Day, &change.Weight, &change.Kcal, &change.CreatedUp, &change.UpdateUp); err != nil {
			return ResponseChange{}, err
		}
		changesData = append(changesData, change)
	}

	return ResponseChange{
		Collection: changesData,
		CollectionTraining: nil,
		Status: http.StatusOK,
		Error: "",
	}, nil
}