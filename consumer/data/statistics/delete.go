package statistics

type Delete struct {
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