package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Auth struct{}


func (p *Auth) Authorization(c *gin.Context) {
	// tokenString := c.GetHeader("Authorization")
	userData := c.GetHeader("UserData")

	
	c.JSON(http.StatusOK, gin.H{"token": userData})
}

func createUser(c *gin.Context){
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"info": err})
	}

	var pk int
	query := `SELECT * FROM users`
	db.QueryRow(query).Scan(pk)




	defer db.Close()

}
func checkUser(){}

