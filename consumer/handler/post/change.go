package post

import (
	"fmt"
	change_data "internal/consumer/data/post"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"internal/consumer/handler/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseChange struct{
	Collection []change_data.Change
	Status     int
	Error      string
}

func HandlerChange(c *gin.Context){
	change, err := change(c)
	if err != nil{
		c.JSON(http.StatusOK, ResponseChange{
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

func change(c* gin.Context)(ResponseChange, error){
	userData := c.GetHeader("UserData")
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

	var changePost change_data.Change
	err = c.BindJSON(&changePost)
	if err != nil{
		return ResponseChange{}, err
	}

	
	id := c.Param("id")


	var updateFields []string
	if changePost.Day != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"day"='%d'`, *changePost.Day))
    }
    if changePost.Weight != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"weight"='%f'`, *changePost.Weight))
    }
    if changePost.Kcal != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"kcal"='%d'`, *changePost.Kcal))
    }
    if changePost.UpdateUp != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"updateUp"='%s'`, *changePost.UpdateUp))
    }
	if changePost.Description != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"description"='%s'`, *changePost.Description))
    }

	if len(updateFields) == 0 {
		if err != nil {
			return ResponseChange{}, err
		}
	}



	query := `UPDATE post SET` +  strings.Join(updateFields, ", ") + ` WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", "day", "weight", "kcal", "createdUp", "updateUp", "description";`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseChange{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var change change_data.Change
		if err := rows.Scan(&change.Id, &change.UserId, &change.Day, &change.Weight, &change.Kcal, &change.CreatedUp, &change.UpdateUp, &change.Description); err != nil {
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