package user

import (
	"fmt"
	"net/http"
	"strings"

	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	auth "myInternal/consumer/handler/auth"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) ChangeUser(c *gin.Context){
	userData := c.GetHeader("UserData")
	var changeBody user_data.User
	var user user_data.User
	var users []user_data.User


	db, err := database.ConnectToDataBase()
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"collection": nil,
			"status": http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	_, users, err = auth.CheckUser(userData)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"collection": nil,
			"status": http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	err = c.BindJSON(&changeBody)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"collection": nil,
			"status": http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	var updateFields []string
	if changeBody.UserName != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"userName"='%s'`, *changeBody.UserName))
    }
    if changeBody.LastName != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"lastName"='%s'`, *changeBody.LastName))
    }
    if changeBody.NickName != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"nickName"='%s'`, *changeBody.NickName))
    }
    if changeBody.Email != nil {
        updateFields = append(updateFields, fmt.Sprintf(`"email"='%s'`, *changeBody.Email))
    }

	if len(updateFields) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"collection": nil,
			"status": http.StatusBadRequest,
			"error":"no fields to update",
		})
		return
	}

	query := `UPDATE users SET ` + strings.Join(updateFields, ", ") + ` WHERE id = $1 RETURNING "id", "userName", "lastName", "nickName", "email", "role";`
	row := db.QueryRow(query, users[0].Id)
	err = row.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role)
	users = []user_data.User{}
	users = append(users, user)
	defer db.Close()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"collection": nil,
			"status": http.StatusBadRequest,
			"error":err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"collection": users,
		"status": http.StatusOK,
		"error":nil,
	})
}
