package post

import (
	delete_data "internal/consumer/data/post"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"internal/consumer/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseDelete struct{
	Collection []delete_data.Delete
	Status     int
	Error      string
}

func HandlerDelete(c *gin.Context){
	delete, err := delete(c)
	if err != nil{
		c.JSON(http.StatusOK, ResponseDelete{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseDelete{
		Collection: delete.Collection,
		Status: delete.Status,
		Error: delete.Error,
	})
}

func delete(c *gin.Context)(ResponseDelete, error){
	userData := c.GetHeader("UserData")
	var usersData []user_data.User
	var deletesData []delete_data.Delete


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDelete{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseDelete{}, err
	}

	usersData = users

	id := c.Param("id")

	query := `DELETE FROM post WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", "day", "weight", "kcal", "createdUp", "updateUp", "description";`
	rows, err := db.Query(query, &id, &usersData[0].Id)
	if err != nil {
		return ResponseDelete{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var delete delete_data.Delete
		if err := rows.Scan(&delete.Id, &delete.UserId, &delete.Day, &delete.Weight, &delete.Kcal, &delete.CreatedUp, &delete.UpdateUp, &delete.Description); err != nil {
			return ResponseDelete{}, err
		}
		deletesData = append(deletesData, delete)
	}

	return ResponseDelete{
		Collection: deletesData,
		Status: http.StatusOK,
		Error: "",
	}, nil

}