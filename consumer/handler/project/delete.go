package project

import (
	params_data "myInternal/consumer/data"
	delete_data "myInternal/consumer/data/project"
	project_data "myInternal/consumer/data/project"
	database "myInternal/consumer/database"
	"myInternal/consumer/handler/auth"
	"net/http"

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
	var deletesData []delete_data.Delete


	db, err := database.ConnectToDataBase()
	if err != nil{
		return ResponseDeleteProject{}, err
	}

	_, _, err = auth.CheckUser(userData)
	if err != nil{
		return ResponseDeleteProject{}, err
	}

	projectId := params.Param


	//TODO AS dorobić usuwanie postów i zdjęć + folder

	query := `DELETE FROM project WHERE "id" = $1 RETURNING "id", "userId", title, description, "createdUp", "updateUp";`
	rows, err := db.Query(query, &projectId)
	if err != nil {
		return ResponseDeleteProject{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var project delete_data.Delete
		if err := rows.Scan(&project.Id, &project.Title, &project.Description, &project.CreatedUp, &project.UpdateUp); err != nil {
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