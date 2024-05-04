package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Post struct{}

func (p *Post) Create(c *gin.Context) {
	fmt.Println("Create Post")

}

func (p *Post) Collection(c *gin.Context) {
	fmt.Println("Collection Post")

}

func (p *Post) GetById(c *gin.Context) {
	fmt.Println("Get Post by Id")

}

func (p *Post) UpdateById(c *gin.Context) {
	fmt.Println("Update Post by Id")

}

func (p *Post) DeleteById(c *gin.Context) {
	fmt.Println("Delete Post by Id")

}
