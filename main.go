package main

import (
	database "internal/consumer/database"
	initializers "internal/consumer/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Collection struct {
	Data []OneResult `json:"data"`
}

type OneResult struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
}


var collection = Collection{
	Data: []OneResult{
		{Id:"1", Name: "artek", Age: "24"},
		{Id:"2", Name: "maks", Age: "24"},
	},
}


func FetchCollection(context *gin.Context){
	context.IndentedJSON(http.StatusOK, collection)
}


func init(){
	initializers.LoadEnv()
	database.ConnectToDataBase()
}


func main(){

	router := gin.Default()
	router.GET("/collection", FetchCollection)
	router.Run("localhost:8080")

}