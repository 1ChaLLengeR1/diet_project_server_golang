package test

import (
	"fmt"
	"mime/multipart"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	project_data "myInternal/consumer/data/project"
	file_function "myInternal/consumer/handler/file"
	project_function "myInternal/consumer/handler/project"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCreateProject(t *testing.T) {

	dataBody := `{
		"title":"test title",
		"description":"desc test"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		AppLanguage: common_test.AppLanguagePL,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	_, err = project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("error create project function: %v", err)
	}
}


func TestCreateFullProject(t *testing.T) {

	dataBody := `{
		"title":"test title",
		"description":"desc test",
		"createdUp":"2024-05-12 10:30:00",
		"updateUp":"2024-05-12 10:30:00"
	}`

	var createProject project_data.Create
	err := helpers.UnmarshalJSONToType(dataBody, &createProject); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ := helpers.BindJSONToMap(&createProject)

	params := params_data.Params{
		Header: common_test.UserTest,
		Param: common_test.TestUUid,
		AppLanguage: common_test.AppLanguagePL,
		Json: jsonMap,
	}

	env.LoadEnv("./.env")
	project, err := project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	formData := make(map[string][]*multipart.FileHeader)
	i := 0
	pathImg := "./consumer/common/test.png"

	fileHeader, file, err := common_test.FileFromPath(pathImg)
	if err != nil {
		t.Fatalf("createFormData: %v", err)
	}
	defer file.Close() 

	formData[fmt.Sprintf("file[%d]", i)] = append(formData[fmt.Sprintf("file[%d]", i)], fileHeader)

	params = params_data.Params{
		Header: common_test.UserTest,
		FormData: formData,
		FormDataParams: map[string]interface{}{
			"projectId": project.Collection[0].Id,
			"folder": *project.Collection[0].Title,
			"names":  []string{"project"},
		},
	}

	_, err = file_function.CreateFile(params)
	if err != nil {
		t.Fatalf("createFile error: %v", err)
	}

}