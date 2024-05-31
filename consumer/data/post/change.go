package post

import training_data "myInternal/consumer/data/training"

type Change struct {
	Id        string   `json:"id"`
	UserId    string   `json:"userId"`
	ProjectId string   `json:"projectId"`
	Day       *int64   `json:"day"`
	Weight    *float64 `json:"weight"`
	Kcal      *int64   `json:"kcal"`
	TrainingCollection *[]training_data.OneTraining `json:"trainingCollection"`
	CreatedUp string   `json:"createdUp"`
	UpdateUp  *string  `json:"updateUp"`
}