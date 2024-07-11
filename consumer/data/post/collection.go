package post

type Collection struct {
	Id        string  `json:"id"`
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Day       int64   `json:"day"`
	Weight    float64 `json:"weight"`
	Kcal      int64   `json:"kcal"`
	CreatedUp string  `json:"createdUp"`
	UpdateUp  string  `json:"updateUp"`
}

// id is userId
type SearchPost struct {
	Id         string `json:"id"`
	ProjectId  string `json:"projectId"`
	IdLanguage string `json:"idLanguage"`
	Page       string `json:"page"`
}