package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	createF "myInternal/consumer/handler/post"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCreatePost(t *testing.T) {

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500,
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00",
		"description":"desc"
	}`

	var createPost post_data.Post
	err := helpers.UnmarshalJSONToType(dataBody, &createPost); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params := params_data.Params{
		Header: common_test.UserDiaxMen,
		Json: jsonMap,
	}

	env.LoadEnv("./../../../../.env")
	_, err = createF.Create(params)
	if err != nil {
		t.Fatalf("Error create function: %v", err)
	}
}