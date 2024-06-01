package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	typeTraining_data "myInternal/consumer/data/typeTraining"
	typeTraining_function "myInternal/consumer/handler/typeTraining"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollectionTypeTraining(t *testing.T) {
	dataBody := `{
		"name":"gym"
	}`

	var createTypeTraining typeTraining_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createTypeTraining); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}

	jsonMap, _ := helpers.BindJSONToMap(&createTypeTraining)

	params := params_data.Params{
		Header: common_test.UserTest,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	_, err = typeTraining_function.CreateTypeTraining(params)
	if err != nil {
		t.Fatalf("Error create typeTraining function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
	}

	collectionTypeTraining, err := typeTraining_function.CollectionTypeTraining(params)
	if err != nil {
		t.Fatalf("Error collection typeTraining function: %v", err)
	}

	if len(collectionTypeTraining.Collection) == 0{
		t.Fatalf("Error collection from typeTraining function is 0: %v", err)
	}
}