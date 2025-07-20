package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) {
	events := models.GetEvents()
	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Events retrieved successfully",
		"data": events,
	})
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  false,
			"message": "Invalid input data",
			"data":    nil,
		})
		return
	}

	event.ID = 1
	event.UserID = 1 // Assuming a static UserID for simplicity

	event.Store()
	context.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Event created successfully",
		"data":    event,
	})
}
