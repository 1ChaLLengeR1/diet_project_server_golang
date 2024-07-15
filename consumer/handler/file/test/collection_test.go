package test

import (
	"fmt"
	"mime/multipart"
	common_test "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	file_function "myInternal/consumer/handler/file"
	common_project "myInternal/consumer/handler/project/test"
	helpers "myInternal/consumer/helper"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCollectionFile(t *testing.T) {
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
		Header:   common_test.UserTest,
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
		Param: createFile.Collection[0].ProjectId,
	}

	collectionFile, err := file_function.FileCollection(params)
	if err != nil {
		t.Fatalf("createFile error: %v", err)
	}

	if len(collectionFile.Collection) == 0{
		t.Fatalf("collectionFile len is 0 - error: %v", err)
	}
}

func TestCollectionMultiple(t *testing.T){
	var params params_data.Params
	formData := make(map[string][]*multipart.FileHeader)
	env.LoadEnv("./.env")

	newProjectId_1, _ := common_project.CreateProject()

	newFile := CreateFile(2)
	formData = newFile

	params = params_data.Params{
		Header:   common_test.UserTest,
		FormData: formData,
		FormDataParams: map[string]interface{}{
			"projectId": newProjectId_1,
			"folder": "testFolder",
			"names":  []string{"test", "test2"},
		},
		
	}

	_, err := file_function.CreateFile(params)
	if err != nil {
		t.Fatalf("createFile error: %v", err)
	}

	ids := file_data.CollectionIds{
		Ids: []string{newProjectId_1},
	}
	
	jsonMap, _ := helpers.BindJSONToMap(&ids)

	params = params_data.Params{
		Json: jsonMap,
	}

	fileCollectionMultiple, err := file_function.FileCollectionMultiple(params)
	if err != nil{
		t.Fatalf("fileCollectionMultiple error: %v", err)
	}

	if len(fileCollectionMultiple.Collection) == 0{
		t.Fatalf("fileCollectionMultiple collection length is 0: %v", err)
	}
}

func TestDownloadProject(t *testing.T){
	env.LoadEnv("./.env")

	params := params_data.Params{
		Param: "5a3f3903-ebf5-4a2c-bf1b-7aab3c5f1681",
	}

	fileZip, err := file_function.CreateZip(params)
	if err != nil {
		t.Fatalf("download project error: %v", err)
	}
	if fileZip == ""{
		t.Fatalf("Not get path to zip!")
	}
}