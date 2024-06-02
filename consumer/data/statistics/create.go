package statistics

type Create struct {
	ID           string  `json:"id"`
	ProjectID    string  `json:"projectId"`
	Day          string  `json:"day"`
	EndWeight    float64 `json:"endWeight"`
	DownWeight   float64 `json:"downWeight"`
	SumKg        float64 `json:"sumKg"`
	AvgKg        float64 `json:"avgKg"`
	SumKcal      int     `json:"sumKcal"`
	TypeTraining string  `json:"typeTraining"`
	SumTime      string  `json:"sumTime"`
	CreatedUp    string  `json:"createdUp"`
	UpdateUp     string  `json:"updateUp"`
}

type OneTraining struct {
	PostId string `json:"postId"`
	Type   string `json:"type"`
	Time   string `json:"time"`
	Kcal   int64  `json:"kcal"`
}

type Statistics struct {
	Day                int64         `json:"day"`
	Weight             float64       `json:"weight"`
	Kcal               int64         `json:"kcal"`
	TrainingCollection []OneTraining `json:"trainingCollection"`
}