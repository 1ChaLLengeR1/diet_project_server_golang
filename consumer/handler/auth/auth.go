package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	user_data "internal/consumer/data/user"
	database "internal/consumer/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct{}


func (p *Auth) Authorization(c *gin.Context) {
	userData := c.GetHeader("UserData")
	var usersData []user_data.User
	
	value, users,  err := checkUser(userData)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	usersData = users

	if value{
		value, err := createUser(userData)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		usersData = value
	}
	
	c.JSON(http.StatusOK, gin.H{
		"collection":usersData,
		"status": http.StatusOK,
	})
}

func checkUser(userData string)(bool, []user_data.User, error){
	var data user_data.UserData
	var user user_data.User
	var users []user_data.User

	db, err := database.ConnectToDataBase()
	if err != nil{
		return false, nil, err
	}

	err = json.Unmarshal([]byte(userData), &data)
	if err != nil {
		return false, nil, fmt.Errorf("error josn userData: %v", err)
	}

	query := `SELECT * FROM users WHERE "email" = $1 AND "nickName" = $2`
	row := db.QueryRow(query, data.Name, data.Nickname)
	err = row.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role)
	users = append(users, user)
	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil, nil
		}
		return false, nil, fmt.Errorf("error scanning row: %v", err)
	}
	
	return false, users, nil
}


func createUser(userData string) ([]user_data.User, error){
	
	db, err := database.ConnectToDataBase()
	if err != nil{
		return nil, err
	}

	var data user_data.UserData
	err = json.Unmarshal([]byte(userData), &data)
	if err != nil {
		return nil, fmt.Errorf("error josn userData: %v", err)
	}

	var users []user_data.User
	query := `INSERT INTO users ("userName", "lastName", "nickName", "email", "role") VALUES ($1, $2, $3, $4, $5) RETURNING "id", "userName", "lastName", "nickName", "email", "role";`

	rows, err := db.Query(query, "", "", data.Nickname, data.Name, "user")
	if err != nil {
		return nil, fmt.Errorf("error db.query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user user_data.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role); err != nil {
			return nil, fmt.Errorf("error scanning user row: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}


