package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Collection struct{}

func (p *Collection) Collection(c *gin.Context) {
	fmt.Println("Collection Post")
}

func (p *Collection) GetById(c *gin.Context) {
	fmt.Println("Get Post by Id")
}