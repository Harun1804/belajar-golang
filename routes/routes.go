package routes

import (
	"example.com/rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", controllers.GetEvents)
	server.GET("/events/:id", controllers.GetEvent)
	server.POST("/events", controllers.CreateEvent)
	server.PUT("/events/:id", controllers.UpdateEvent)
}