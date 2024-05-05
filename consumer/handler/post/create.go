package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Create struct{}

func (p *Create) Create(c *gin.Context) {
	fmt.Println("Create Post")

}