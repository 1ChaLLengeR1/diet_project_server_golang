package file

type Create struct {
	Id        string `json:"id"`
	PostId    string `json:"postId"`
	Name      string `json:"name"`
	Folder    string `json:"folder"`
	Path      string `json:"path"`
	Url       string `json:"url"`
	CreatedUp string `json:"createdUp"`
	UpdateUp  string `json:"updateUp"`
}