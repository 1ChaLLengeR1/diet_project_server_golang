package post

type Collection struct {
	Id          string  `json:"id"`
	UserId      string  `json:"userId"`
	Day         int64   `json:"day"`
	Weight      float64 `json:"weight"`
	Kcal        int64   `json:"kcal"`
	CreatedUp   string  `json:"createdUp"`
	UpdateUp    string  `json:"updateUp"`
	Description string  `json:"description"`
}