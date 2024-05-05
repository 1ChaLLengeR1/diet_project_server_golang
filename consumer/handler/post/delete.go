package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Delete struct{}

func (p *Delete) DeleteById(c *gin.Context) {
	fmt.Println("Delete Post by Id")

}
