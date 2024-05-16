package file

import (
	"fmt"
	"io"
	"mime/multipart"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	random "myInternal/consumer/helper"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


type ResponseFileCreate struct {
	Collection []file_data.Create 	`json:"collection"`
	Status     int 					`json:"status"`
	Error      string 				`json:"error"`
}

func HandlerCreateFile(ctx *gin.Context) {
    formData := make(map[string][]*multipart.FileHeader)
    var nameData []string

    
    for i := 0; ; i++ {
        file, err := ctx.FormFile(fmt.Sprintf("file[%d]", i))
        if err != nil {
            if err == http.ErrMissingFile {
                break 
            }
            ctx.JSON(http.StatusBadRequest, ResponseFileCreate{
                Collection: nil,
                Status:     http.StatusBadRequest,
                Error:      err.Error(),
            })
            return
        }
        formData[fmt.Sprintf("file[%d]", i)] = append(formData[fmt.Sprintf("file[%d]", i)], file)
    }

    
    for j := 0; ; j++ {
        name := ctx.PostForm(fmt.Sprintf("name[%d]", j))
        if name == "" {
            break
        }
        nameData = append(nameData, name)
    }

    
    params := params_data.Params{
        Header: ctx.GetHeader("UserData"),
        FormData: formData,
        FormDataParams: map[string]interface{}{
            "postId": ctx.PostForm("postId"),
            "folder": ctx.PostForm("folder"),
			"names": nameData,
        },
    }

    
    _, err := CreateFile(params)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, ResponseFileCreate{
            Collection: nil,
            Status:     http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }
}

func CreateFile(params params_data.Params)(ResponseFileCreate, error){
	userData := params.Header
	var usersData []user_data.User
	var uploadedFiles []string

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileCreate{}, err
	}

	_, users,  err := auth.CheckUser(userData)
	if err != nil{
		return ResponseFileCreate{}, err
	}

	usersData = users
    index := 0

	for _, files := range params.FormData {
        for _, file := range files {
            src, err := file.Open()
            if err != nil {
                return ResponseFileCreate{}, err
            }
            defer src.Close()

			fileExtension := filepath.Ext(file.Filename)
			fileNameWithoutExt := file.Filename[:len(file.Filename)-len(fileExtension)]
			randomStr, err := random.GenerateRandomString(8)
            if err != nil {
                return ResponseFileCreate{}, err
            }
            
			folder := params.FormDataParams["folder"].(string)
			folderPath := filepath.Join("consumer", "file", "posts", folder)
			if _, err := os.Stat(folderPath); os.IsNotExist(err) {
				if err := os.MkdirAll(folderPath, 0755); err != nil {
					return ResponseFileCreate{}, err
				}
			}

			var fileName string
			if index < len(params.FormDataParams["names"].([]string)) {
				name := params.FormDataParams["names"].([]string)[index]
				fileName = fmt.Sprintf("%s_%s_%s%s", fileNameWithoutExt, name, randomStr, fileExtension)
			} else {
				fileName = fmt.Sprintf("%s_%s%s", fileNameWithoutExt, randomStr, fileExtension)
			}


            dstPath := filepath.Join(folderPath, fileName)
            dst, err := os.Create(dstPath)
            if err != nil {
                return ResponseFileCreate{}, err
            }
            defer dst.Close()

            if _, err := io.Copy(dst, src); err != nil {
                return ResponseFileCreate{}, err
            }
            index++;
        }
        
    }




	
	fmt.Println(uploadedFiles, usersData, db)

	return ResponseFileCreate{}, err
}