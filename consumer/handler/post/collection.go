package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func Collection(c *gin.Context) {
	fmt.Println("Collection Post")
}

func GetById(c *gin.Context) {
	fmt.Println("Get Post by Id")
}