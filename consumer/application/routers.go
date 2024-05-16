package application

import (
	auth_handler "myInternal/consumer/handler/auth"
	file_handler "myInternal/consumer/handler/file"
	post_handler "myInternal/consumer/handler/post"
	user_handler "myInternal/consumer/handler/user"
	"myInternal/consumer/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)


func loadRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "Start routers Gin")
	})

	//post routers
	postGroup := router.Group("/api/post")
	{
		postGroup.POST("/create", middleware.EnsureValidToken(), post_handler.CreateHandler)
		postGroup.GET("/collection/:page", middleware.EnsureValidToken(), post_handler.HandlerCollection)
		postGroup.GET("/one/:id", post_handler.HandlerCollectionOne)
		postGroup.PATCH("/change/:id", middleware.EnsureValidToken(), post_handler.HandlerChange)
		postGroup.DELETE("/delete/:id", middleware.EnsureValidToken(), post_handler.HandlerDelete)
	}

	//file routers
	fileGroup := router.Group("/api/file")
	{
		fileGroup.POST("/create", middleware.EnsureValidToken(), file_handler.HandlerCreateFile)
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


	router.GET("/api/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from a public endpoint! You don't need to be authenticated to see this."})
	})

	// This route is only accessible if the user has a valid access_token.
	router.GET("/api/private", middleware.EnsureValidToken(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from a private endpoint! You need to be authenticated to see this."})
	}) 

	return router
}




 



