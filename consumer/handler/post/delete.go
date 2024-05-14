package post

import (
	params_data "internal/consumer/data"
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

	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Query: c.Query("private"),
		Param: c.Param("id"),
	}


	delete, err := delete(c, params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseDelete{
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

func delete(c *gin.Context, params params_data.Params)(ResponseDelete, error){
	userData := params.Header
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

	id := params.Param

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