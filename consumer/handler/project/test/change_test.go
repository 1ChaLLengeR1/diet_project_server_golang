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

func TestChangeProject(t *testing.T) {
	dataBody := `{
		"title":"test title",
		"description":"desc test",
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject)
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param:  common_test.TestUUid,
		Json:   jsonMap,
	}

	env.LoadEnv("./.env")
	project, err := project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("Error create function: %v", err)
	}

	dataBody = `{
		"title":"test title update",
		"description":"desc test update",
		"updateUp":"2024-06-12 10:00:00"
	}`

	var changeProject project_data.Create
	err = helpers.UnmarshalJSONToType(dataBody, &changeProject)
	if err != nil {
		t.Fatalf("Error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ = helpers.BindJSONToMap(&changeProject)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param:  project.Collection[0].Id,
		Json:   jsonMap,
	}

	_, err = project_function.ChangeProject(params)
	if err != nil{
		t.Fatalf("Error change project function: %v", err)
	}
}