package common

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func FileFromPath(filePath string) (*multipart.FileHeader, *os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, nil, err
	}

	// Create a buffer to write our multipart form data
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Create the form file field
	formFile, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		file.Close()
		return nil, nil, err
	}

	// Copy the file content into the form file field
	_, err = io.Copy(formFile, file)
	if err != nil {
		file.Close()
		return nil, nil, err
	}

	// Close the writer to finalize the form data
	writer.Close()

	// Parse the multipart form data to get the FileHeader
	reader := multipart.NewReader(&b, writer.Boundary())
	form, err := reader.ReadForm(int64(fileInfo.Size()))
	if err != nil {
		file.Close()
		return nil, nil, err
	}

	// Extract the file header from the parsed form data
	fileHeaders := form.File["file"]
	if len(fileHeaders) == 0 {
		file.Close()
		return nil, nil, fmt.Errorf("no file headers found")
	}

	return fileHeaders[0], file, nil
}