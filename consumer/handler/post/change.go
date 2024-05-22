package post

import (
	"fmt"
	params_data "myInternal/consumer/data"
	change_data "myInternal/consumer/data/post"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	helpers "myInternal/consumer/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseChange struct{
	Collection []change_data.Change `json:"collection"`
	Status     int              	`json:"status"`
	Error      string          		`json:"error"`
}



func HandlerChange(c *gin.Context){

	var changePost change_data.Change
	c.BindJSON(&changePost)
	jsonMap, err := helpers.BindJSONToMap(&changePost)
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
		Param: c.Param("id"),
		Json: jsonMap,
	}


	change, err := Change(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseChange{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseChange{
		Collection: change.Collection,
		Status: change.Status,
		Error: change.Error,
	})
}

func Change(params params_data.Params)(ResponseChange, error){
	userData := params.Header
	
	var usersData []user_data.User
	var changesData []change_data.Change


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseChange{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseChange{}, err
	}

	usersData = users

	id := params.Param

	day, dayOk := params.Json["day"].(float64) 
	weight, weightOk := params.Json["weight"].(float64)
	kcal, kcalOk := params.Json["kcal"].(float64)
	updateUp, updateUpOk := params.Json["updateUp"].(string)
	description, descriptionOk := params.Json["description"].(string)

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
	if updateUpOk {
		updateFields = append(updateFields, fmt.Sprintf(`"updateUp"='%s'`, updateUp))
	}
	if descriptionOk {
		updateFields = append(updateFields, fmt.Sprintf(`"description"='%s'`, description))
	}

	if len(updateFields) == 0 {
		if err != nil {
			return ResponseChange{}, err
		}
	}

	query := `UPDATE post SET` +  strings.Join(updateFields, ", ") + ` WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", "projectId", "day", "weight", "kcal", "createdUp", "updateUp", "description";`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseChange{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var change change_data.Change
		if err := rows.Scan(&change.Id, &change.UserId, &change.ProjectId, &change.Day, &change.Weight, &change.Kcal, &change.CreatedUp, &change.UpdateUp, &change.Description); err != nil {
			return ResponseChange{}, err
		}
		changesData = append(changesData, change)
	}

	return ResponseChange{
		Collection: changesData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}