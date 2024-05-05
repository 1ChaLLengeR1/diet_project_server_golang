package post

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Change struct{}

func (p *Change) UpdateById(c *gin.Context) {
	fmt.Println("Update Post by Id")
}
