package application

import (
	auth_handler "internal/consumer/handler/auth"
	post_handler "internal/consumer/handler/post"
	user_handler "internal/consumer/handler/user"
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
	postGroup := router.Group("api/post")
	{
		postGroup.POST("/create", post_handler.CreateHandler)
		// postGroup.GET("/collection", post_handler.)
		// postGroup.GET("/getById/:id", post_handler.GetById)
		postGroup.PATCH("/change/:id", post_handler.HandlerChange)
		postGroup.DELETE("/delete/:id", post_handler.HandlerDelete)
	}

	// auth jwt
	authHander := &auth_handler.Auth{}
	authGroup := router.Group("/api/auth", middleware.EnsureValidToken())
	{
		authGroup.POST("/authorization", authHander.Authorization)
	}

	//user
	userHandler :=  &user_handler.User{}
	userGroup := router.Group("/api/user", middleware.EnsureValidToken())
	{
		userGroup.PATCH("/change", userHandler.ChangeUser)
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




 



