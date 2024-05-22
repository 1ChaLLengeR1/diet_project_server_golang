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

func TestCollectionAll(t *testing.T) {
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
	err = helpers.UnmarshalJSONToType(dataBody, &createPost)
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")

	for i := 0; i < 3; i++ {
		_, err := post_function.Create(params)
		if err != nil {
			t.Fatalf("Error in create function: %v", err)
		}
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: "1",
		Json: jsonMap,
	}

	valueCollection, err := post_function.Collection(params)
	if err != nil {
		t.Fatalf("Error in collection function: %v", err)
	}

	if(len(valueCollection.Collection) < 3){
		t.Fatalf("Error in len collection > 3")
	}
}


func TestCollectionOne(t *testing.T){
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
	err = helpers.UnmarshalJSONToType(dataBody, &createPost)
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json:   jsonMap,
	}

	env.LoadEnv("./.env")
	valueCreate, err := post_function.Create(params)
	if err != nil {
			t.Fatalf("Error in create function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Query: "true",
		Param: valueCreate.Collection[0].Id,
	}

	_, err = post_function.CollectionOne(params)
	if err != nil {
		t.Fatalf("Error in create function: %v", err)
	}

}