package application

import (
	auth_handler "myInternal/consumer/handler/auth"
	file_handler "myInternal/consumer/handler/file"
	post_handler "myInternal/consumer/handler/post"
	project_handler "myInternal/consumer/handler/project"
	user_handler "myInternal/consumer/handler/user"
	"myInternal/consumer/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)


func loadRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Static("/consumer/file", "./consumer/file")

	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "Start routers Gin")
	})

	// project routers

	projectGroup := router.Group("/api/project")
	{
		projectGroup.POST("/create",  middleware.EnsureValidToken(), project_handler.HandlerCreateProject)
	}

	//post routers
	postGroup := router.Group("/api/post")
	{
		postGroup.POST("/create/:projectId", middleware.EnsureValidToken(), post_handler.CreateHandler)
		postGroup.GET("/collection/:page", middleware.EnsureValidToken(), post_handler.HandlerCollection)
		postGroup.GET("/one/:id", post_handler.HandlerCollectionOne)
		postGroup.PATCH("/change/:id", middleware.EnsureValidToken(), post_handler.HandlerChange)
		postGroup.DELETE("/delete/:id", middleware.EnsureValidToken(), post_handler.HandlerDelete)
	}

	//file routers
	fileGroup := router.Group("/api/file")
	{
		fileGroup.POST("/create", middleware.EnsureValidToken(), file_handler.HandlerCreateFile)
		fileGroup.DELETE("/delete/:deleteId", middleware.EnsureValidToken(), file_handler.HandlerFileDelete)
		fileGroup.GET("/collection/:postId", file_handler.HandlerFileCollection)
	}

	// auth jwt
	authHander := &auth_handler.Auth{}
	authGroup := router.Group("/api/auth", middleware.EnsureValidToken())
	{
		authGroup.POST("/authorization", authHander.Authorization)
	}

	//user routers
	userGroup := router.Group("/api/user", middleware.EnsureValidToken())
	{
		userGroup.PATCH("/change", user_handler.HandlerChangeUser)
	}

	return router
}




 



