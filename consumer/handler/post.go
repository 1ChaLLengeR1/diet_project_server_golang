package handler

import (
	"fmt"
	"net/http"
)

type Post struct{}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Post")
}

func (p *Post) Collection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Collection Post")
}

func (p *Post) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Post by Id")
}

func (p *Post) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Post by Id")
}

func (p *Post) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Post by Id")
}