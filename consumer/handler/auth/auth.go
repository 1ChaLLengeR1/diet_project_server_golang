package auth

import (
	"encoding/json"
	"fmt"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct{}


func (p *Auth) Authorization(c *gin.Context) {
	// tokenString := c.GetHeader("Authorization")
	userData := c.GetHeader("UserData")
	// err := createUser(c)
	// if err != nil{
	// 	c.JSON(http.StatusBadRequest, gin.H{"error: ": err})
	// 	return
	// }

	err := checkUser(userData)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"userData": userData})
}

func checkUser(userData string)error{
	var data user_data.UserData
	var user user_data.User

	db, err := database.ConnectToDataBase()
	if err != nil{
		return err
	}

	err = json.Unmarshal([]byte(userData), &data)
	if err != nil {
		return fmt.Errorf("error josn userData: %v", err)
	}

	query := `SELECT * FROM users WHERE "userName" = $1 AND "nickName" = $2`

	var users []user_data.User
	
	row := db.QueryRow(query, data.Name, data.Nickname)
	err = row.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role)

	users = append(users, user)
	defer db.Close()

	if err != nil{
		return fmt.Errorf("error queryRow: %v", err)
	}
	fmt.Println(users)
	
	return nil
}


func createUser(c *gin.Context) error{
	
	db, err := database.ConnectToDataBase()
	if err != nil{
		return err
	}

	var users []user_data.User


	query := `SELECT * FROM users;`
	rows, err := db.Query(query)
	if err != nil{
		return fmt.Errorf("error db.query: %v", err)
	}
	defer rows.Close()

	for  rows.Next(){
		var user user_data.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role); err != nil {
			return fmt.Errorf("error rows.scan: %v", err)
		}
		users = append(users, user)
	}

	fmt.Println(users)
	defer db.Close()
	return nil

}


