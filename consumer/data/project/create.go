package project

type Create struct {
	Id          string  `json:"id"`
	UserId      string  `json:"userId"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	CreatedUp   *string `json:"createdUp"`
	UpdateUp    *string `json:"updateUp"`
}