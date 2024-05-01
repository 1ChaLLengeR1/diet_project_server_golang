package application

import (
	"internal/consumer/handler"
	auth_handler "internal/consumer/handler/auth"
	"internal/consumer/middleware"
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
	postHandler := &handler.Post{}
	postGroup := router.Group("/post")
	{
		postGroup.POST("/", postHandler.Create)
		postGroup.GET("/", postHandler.Collection)
		postGroup.GET("/:id", postHandler.GetById)
		postGroup.PATCH("/:id", postHandler.UpdateById)
		postGroup.DELETE("/:id", postHandler.DeleteById)
	}

	// auth jwt
	authHander := &auth_handler.Auth{}
	authGroup := router.Group("/api/auth", middleware.EnsureValidToken())
	{
		authGroup.POST("/authorization", authHander.Authorization)
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




 



