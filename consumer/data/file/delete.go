package file

type Delete struct {
	Id        string `json:"id"`
	ProjectId string `json:"projectId"`
	Name      string `json:"name"`
	Folder    string `json:"folder"`
	Path      string `json:"path"`
	Url       string `json:"url"`
	CreatedUp string `json:"createdUp"`
	UpdateUp  string `json:"updateUp"`
}