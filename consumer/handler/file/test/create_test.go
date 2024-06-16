package test

import (
	"fmt"
	"mime/multipart"
	common_test_params "myInternal/consumer/common"
	helpers "myInternal/consumer/common"
	params_data "myInternal/consumer/data"
	file_function "myInternal/consumer/handler/file"
	env "myInternal/consumer/initializers"
	"testing"
)

func TestCreateFile(t *testing.T) {
	var params params_data.Params
	formData := make(map[string][]*multipart.FileHeader)
	i := 0
	pathImg := "./consumer/common/test.png"

	fileHeader, file, err := helpers.FileFromPath(pathImg)
	if err != nil {
		t.Fatalf("createFormData: %v", err)
	}
	defer file.Close() 

	formData[fmt.Sprintf("file[%d]", i)] = append(formData[fmt.Sprintf("file[%d]", i)], fileHeader)

	params = params_data.Params{
		Header: common_test_params.UserTest,
		FormData: formData,
		FormDataParams: map[string]interface{}{
			"projectId": common_test_params.TestUUid,
			"folder": "testFolder",
			"names":  []string{"test"},
		},
	}

	env.LoadEnv("./.env")
	_, err = file_function.CreateFile(params)
	if err != nil {
		t.Fatalf("createFile error: %v", err)
	}
}
