package main

import (
	"internal/database"
	"internal/initializers"

	"github.com/gin-gonic/gin"
)




func init(){
	initializers.LoadEnv()
	database.ConnectToDataBase()
}


func main(){

	router := gin.Default()
	router.Run("localhost:8080")
}