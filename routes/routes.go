package routes

import (
	"example.com/rest-api/controllers"
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", controllers.CreateEvent)
	authenticated.PUT("/events/:id", controllers.UpdateEvent)
	authenticated.DELETE("/events/:id", controllers.DeleteEvent)
	authenticated.POST("/events/:id/register", controllers.RegisterEvent)
	authenticated.DELETE("/events/:id/cancel", controllers.CancelEvent)

	server.POST("/signup", controllers.StoreUser)
	server.POST("/signin", controllers.LoginUser)
}