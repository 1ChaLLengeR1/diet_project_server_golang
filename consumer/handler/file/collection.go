package file

import (
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseFileCollection struct {
	Collection []file_data.Collection 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}

func HandlerFileCollection(c *gin.Context){
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
	}

	fileCollection, err := FileCollection(params)
	if err != nil {
        c.JSON(http.StatusBadRequest, ResponseFileDelete{
            Collection: nil,
            Status:     http.StatusBadRequest,
            Error:      err.Error(),
        })
        return
    }

	c.JSON(http.StatusOK, ResponseFileCollection{
		Collection: fileCollection.Collection,
		Status: fileCollection.Status,
		Error: fileCollection.Error,
	})
}

func FileCollection(params params_data.Params)(ResponseFileCollection, error){
	userData := params.Header
    var filesData []file_data.Collection

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseFileCollection{}, err
	}

	_, _,  err = auth.CheckUser(userData)
	if err != nil{
		return ResponseFileCollection{}, err
	}

	projectId := params.Param

	query := `SELECT * FROM images WHERE "projectId" = $1`
	rows, err := db.Query(query, &projectId)
	if err != nil {
		return ResponseFileCollection{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var collection file_data.Collection
		if err := rows.Scan(&collection.Id, &collection.ProjectId, &collection.Name, &collection.Folder, &collection.FolderPath, &collection.Path, &collection.Url, &collection.CreatedUp, &collection.UpdateUp); err != nil {
			return ResponseFileCollection{}, err
		}
		filesData = append(filesData, collection)
	}

	return ResponseFileCollection{
		Collection: filesData,
		Status: http.StatusOK,
		Error: "",
	}, nil

}