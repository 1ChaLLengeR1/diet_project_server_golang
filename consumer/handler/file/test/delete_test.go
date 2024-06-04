package test

import (
	"fmt"
	"mime/multipart"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	post_data "myInternal/consumer/data/post"
	project_data "myInternal/consumer/data/project"
	file_function "myInternal/consumer/handler/file"
	post_function "myInternal/consumer/handler/post"
	project_function "myInternal/consumer/handler/project"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	var params params_data.Params
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
			"projectId": common_test.TestUUid,
			"folder": "testFolder",
			"names":  []string{"test"},
		},
	}

	env.LoadEnv("./.env")
	createFile, err := file_function.CreateFile(params)
	if err != nil {
		t.Fatalf("createFile error: %v", err)
	}


	params = params_data.Params{
		Header: common_test.UserTest,
		Param: createFile.Collection[0].Id,
	}

	_, err = file_function.DeleteFile(params)
	if err != nil {
		t.Fatalf("deleteFile error: %v", err)
	}
}

func TestDeleteAll(t *testing.T){
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
	project, err := project_function.CreateProject(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	// post function
	dataBody = `{
		"day":1,
		"weight":88,
		"kcal":2500,
		"description":"desc"
	}`

	var createPost post_data.Post
	err = helpers.UnmarshalJSONToType(dataBody, &createPost); 
	if err != nil {
		t.Fatalf("error unmarshalling dataBody: %v", err)
	}
	jsonMap, _ = helpers.BindJSONToMap(&createPost)

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: project.Collection[0].Id,
		Json: jsonMap,
	}

	_, err = post_function.Create(params)
	if err != nil {
		t.Fatalf("error create function: %v", err)
	}

	// file function
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
			"folder": "testFolder",
			"names":  []string{"test"},
		},
	}

	_, err = file_function.CreateFile(params)
	if err != nil {
		t.Fatalf("createFile error: %v", err)
	}

	params = params_data.Params{
		Header: common_test.UserTest,
		Param: project.Collection[0].Id,
	}

	// delete project
	deleteProject, err := project_function.DeleteProject(params)
	if err != nil {
		t.Fatalf("error delete project function: %v", err)
	}
	
	removeIds := file_data.RemoveId{
		Ids: deleteProject.CollectionRemoveId,
	}
	
	jsonMap, _ = helpers.BindJSONToMap(&removeIds)

	params = params_data.Params{
		Header: common_test.UserTest,
		Json: jsonMap,
	}

	_, err = file_function.DeleteFileAll(params)
	if err != nil {
		t.Fatalf("error deleteAll file function: %v", err)
	}
}