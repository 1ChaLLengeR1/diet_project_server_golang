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
		"kcal":2500
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	valueCreate, err := post_function.Create(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}


	var changePost post_data.Change

	dataChangeBody := `{
		"day":100,
		"weight":88.5,
		"kcal":3500
	}`

	err = helpers.UnmarshalJSONToType(dataChangeBody, &changePost); 
	if err != nil {
		t.Fatalf("error unmarshalling dataChangeBody: %v", err)
	}

	jsonMap, _ = helpers.BindJSONToMap(&changePost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: valueCreate.Collection[0].Id,
		Json: jsonMap,
	}

	_, err = post_function.Change(params)
	if err != nil {
		t.Fatalf("error change function: %v", err)
	}

}