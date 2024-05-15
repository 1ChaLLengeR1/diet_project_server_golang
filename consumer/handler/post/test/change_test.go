package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestChangeAll(t *testing.T) {

	var err error
	var params params_data.Params

	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500,
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00",
		"description":"desc"
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserDiaxMen,
		Json: jsonMap,
	}

	env.LoadEnv("./../../../../.env")
	valueCreate, err := post_function.Create(params)
	if err != nil {
		t.Fatalf("Error create function: %v", err)
	}


	var changePost post_data.Change

	dataChangeBody := `{
		"day":100,
		"weight":88.5,
		"kcal":3500,
		"updateUp":"2025-05-12T10:30:00+02:00",
		"description":"update desssssc 123 123"
	}`

	err = helpers.UnmarshalJSONToType(dataChangeBody, &changePost); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataChangeBody: %v", err)
	}

	jsonMap, _ = helpers.BindJSONToMap(&changePost)

	params = params_data.Params{
		Header: common_test.UserDiaxMen,
		Param: valueCreate.Collection[0].Id,
		Json: jsonMap,
	}

	_, err = post_function.Change(params)
	if err != nil {
		t.Fatalf("Error change function: %v", err)
	}

}