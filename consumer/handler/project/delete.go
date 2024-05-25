package project

import (
	"fmt"
	params_data "myInternal/consumer/data"
	file_data "myInternal/consumer/data/file"
	post_data "myInternal/consumer/data/post"
	project_data "myInternal/consumer/data/project"
	user_data "myInternal/consumer/data/user"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type ResponseDeleteProject struct {
	Collection []project_data.Delete 	`json:"collection"`
	Status     int 						`json:"status"`
	Error      string 					`json:"error"`
}


func HandlerDeleteProject(c *gin.Context) {
	params := params_data.Params{
		Header: c.GetHeader("UserData"),
		Param: c.Param("projectId"),
	}

	projectDelete, err := DeleteProject(params)
	if err != nil{
		c.JSON(http.StatusBadRequest, ResponseDeleteProject{
			Collection: nil,
			Status: http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseDeleteProject{
		Collection: projectDelete.Collection,
		Status: projectDelete.Status,
		Error: projectDelete.Error,
	})
}

func DeleteProject(params params_data.Params)(ResponseDeleteProject, error){
	userData := params.Header
	var usersData []user_data.User
	var deletesData []project_data.Delete
	var folderRemove []string

	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDeleteProject{}, err
	}

	_, users, err := auth.CheckUser(userData)
	if err != nil{
		return ResponseDeleteProject{}, err
	}

	usersData = users
	projectId := params.Param
	folderRemove = append(folderRemove, fmt.Sprintf("'%s'", projectId))

	query := `DELETE FROM post WHERE "projectId" = $1 AND "userId" = $2 RETURNING "id", "userId", "projectId", "day", "weight", "kcal", "createdUp", "updateUp", "description";`
	rows, err := db.Query(query, &projectId, usersData[0].Id)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()

	var delete post_data.Delete
	for rows.Next() {
		if err := rows.Scan(&delete.Id, &delete.UserId, &delete.ProjectId, &delete.Day, &delete.Weight, &delete.Kcal, &delete.CreatedUp, &delete.UpdateUp, &delete.Description); err != nil {
			return ResponseDeleteProject{}, err
		}
		folderRemove = append(folderRemove, fmt.Sprintf("'%s'", delete.Id))
	}

	query = `DELETE FROM images WHERE "projectId" IN (` + strings.Join(folderRemove, ", ") + `) RETURNING "id", "projectId", "name", "folder", "folderPath", "path", "url", "createdUp", "updateUp";`
	rows, err = db.Query(query)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()
	folderRemove = []string{}

	var file file_data.Delete
	for rows.Next(){
		if err := rows.Scan(&file.Id, &file.ProjectId, &file.Name, &file.Folder, &file.FolderPath,  &file.Path, &file.Url, &file.CreatedUp, &file.UpdateUp); err != nil {
			return ResponseDeleteProject{}, err
		}
		folderRemove = append(folderRemove, file.FolderPath)
	}

	err = removeFolders(folderRemove)
    if err != nil {
		return ResponseDeleteProject{}, err
    }
	
	query = `DELETE FROM project WHERE "id" = $1 AND "userId" = $2 RETURNING "id", "userId", title, description, "createdUp", "updateUp";`
	rows, err = db.Query(query, &projectId, usersData[0].Id)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var project project_data.Delete
		if err := rows.Scan(&project.Id, &project.UserId, &project.Title, &project.Description, &project.CreatedUp, &project.UpdateUp); err != nil {
			return ResponseDeleteProject{}, err
		}
		deletesData = append(deletesData, project)
	}

	return ResponseDeleteProject{
		Collection: deletesData,
		Status: http.StatusOK,
		Error: "",
	}, nil
}


func removeFolders(paths []string) error {
    for _, path := range paths {
        err := os.RemoveAll(path)
        if err != nil {
            return err
        }
    }
    return nil
}