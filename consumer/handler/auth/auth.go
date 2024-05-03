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
	
	value, err := checkUser(userData)
	if err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if value{
		value, err := createUser(userData)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(value)
	}
	
	c.JSON(http.StatusOK, gin.H{"success": userData})
}

func checkUser(userData string)(bool, error){
	var data user_data.UserData
	var user user_data.User

	db, err := database.ConnectToDataBase()
	if err != nil{
		return false, err
	}

	err = json.Unmarshal([]byte(userData), &data)
	if err != nil {
		return false, fmt.Errorf("error josn userData: %v", err)
	}

	
	query := `SELECT * FROM users WHERE "email" = $1 AND "nickName" = $2`

	row := db.QueryRow(query, data.Name, data.Nickname)
	err = row.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role)
	defer db.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}
		return false, fmt.Errorf("error scanning row: %v", err)
	}
	
	return false, nil
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
	query := `INSERT INTO users ("userName", "lastName", "nickName", "email", "role") VALUES ($1, $2, $3, $4, $5);`
	rows, err := db.Query(query, "", "",  data.Nickname, data.Name, "user")
	if err != nil{
		return nil, fmt.Errorf("error db.query: %v", err)
	}
	defer rows.Close()

	for  rows.Next(){
		var user user_data.User
		if err := rows.Scan(&user.Id, &user.UserName, &user.LastName, &user.NickName, &user.Email, &user.Role); err != nil {
			return nil, fmt.Errorf("error rows.scan: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}


