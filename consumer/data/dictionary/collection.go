package dictionary

type Collection struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	Translation string `json:"translation"`
	CreatedUp   string `json:"createdUp"`
	UpdateUp    string `json:"updateUp"`
}