package test

import (
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	project_function "myInternal/consumer/handler/project"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollectionProject(t *testing.T) {
	dataBody := `{
		"title":"test title",
		"description":"desc test",
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	_, err = project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("Error create function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: "1",
		Query: "true",
	}

	projectCollection, err := project_function.CollectionProject(params)
	if err != nil {
		t.Fatalf("Error collection function: %v", err)
	}

	if(len(projectCollection.Collection) == 0){
		t.Fatalf("Error in len collection == 0")
	}
}

func TestCollectionOne(t *testing.T){
	dataBody := `{
		"title":"test title",
		"description":"desc test",
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject); 
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	project, err := project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("Error create function: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: project.Collection[0].Id,
		Query: "true",
	}

	projectCollection, err := project_function.CollectionOneProject(params)
	if err != nil {
		t.Fatalf("Error collection one function: %v", err)
	}

	if(len(projectCollection.Collection) == 0){
		t.Fatalf("Error in len collection one == 0")
	}
}