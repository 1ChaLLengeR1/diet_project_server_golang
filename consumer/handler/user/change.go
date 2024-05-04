package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Change(c *gin.Context) {




	c.JSON(http.StatusOK, gin.H{
		"collection": nil,
		"status": http.StatusOK,
	})
}
