package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	post_data "myInternal/consumer/data/post"
	post_function "myInternal/consumer/handler/post"
	training_function "myInternal/consumer/handler/training"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollectioneOne(t *testing.T) {
	dataBody := `{
		"day":1,
		"weight":88,
		"kcal":2500
	}`

	var createPost post_data.Post
	err := helpers.UnmarshalJSONToType(dataBody, &createPost); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createPost)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	postCreate, err := post_function.Create(params)
	if err != nil {
		t.Fatalf("Error create post function: %v", err)
	}

	trainingCollection := `
	{
		"trainingCollection": [
			{
				"type":"gym",
				"time":"2:05:32",
				"kcal":986
			},
			{
				"type":"bike",
				"time":"00:50:19",
				"kcal":543
			},
			{
				"type":"bike",
				"time":"00:48:21",
				"kcal":491
			}
		]
	}
	`
	
	var trainingCollectionMap map[string]interface{}
	err = helpers.UnmarshalJSONToType(trainingCollection, &trainingCollectionMap)
	if err != nil {
		t.Fatalf("Error unmarshalling trainingCollection: %v", err)
	}
	
	jsonMap, err = helpers.BindJSONToMap(&trainingCollectionMap)
	if err != nil {
		t.Fatalf("Error binding JSON to map array: %v", err)
	}
	
	params = params_data.Params{
		Param: postCreate.Collection[0].Id,
		Json:  jsonMap,
	}

	_, err = training_function.CreateTraining(params)
	if err != nil {
		t.Fatalf("Error create training function: %v", err)
	}

	params = params_data.Params{
		Param: postCreate.Collection[0].Id,
	}

	collectionOne, err := training_function.CollectionOneTraining(params)
	if err != nil {
		t.Fatalf("Error collectionOneTraining function: %v", err)
	}
	if(len(collectionOne.Collection) == 0){
		t.Fatalf("Error collection from traning function is 0: %v", err)
	}
}