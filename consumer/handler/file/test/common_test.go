package test

import (
	"fmt"
	"mime/multipart"
	helpers "myInternal/consumer/common"
)

func CreateFile(iterable int) map[string][]*multipart.FileHeader {
	formData := make(map[string][]*multipart.FileHeader)

	var pathImg [2]string
	pathImg[0] = "./consumer/common/test.png"
	pathImg[1] = "./consumer/common/test1.png"

	for i := 0; i < iterable; i++ {
		fileHeader, file, _ := helpers.FileFromPath(pathImg[i])
		defer file.Close()
		formData[fmt.Sprintf("file[%d]", i)] = append(formData[fmt.Sprintf("file[%d]", i)], fileHeader)
	}

	return formData
}