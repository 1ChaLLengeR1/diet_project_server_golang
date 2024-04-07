package application

import (
	"encoding/gob"
	"internal/consumer/authenticator"
	"internal/consumer/handler"
	"log"
	"net/http"

	auth_routers "internal/consumer/web"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)


func loadRouters() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())

	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "Start routers Gin")
	})

	
	// auth routers
	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "consumer/web/static")
	router.LoadHTMLGlob("consumer/web/template/*")

	router.GET("/", func(ctx *gin.Context) { 
		ctx.HTML(http.StatusOK, "home.html", nil)
	})
	router.GET("/login", auth_routers.HandlerLogin(auth))
	router.GET("/callback", auth_routers.HandlerCallback(auth))
	router.GET("/logout", auth_routers.HandlerLogout)
	router.GET("/user", auth_routers.IsAuthenticated, auth_routers.HandlerUser)

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

	return router
}








