package controllers

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to retrieve events",
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Events retrieved successfully",
		"data": events,
	})
}

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid event ID",
			"data":    nil,
		})
		return
	}
	
	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status":  false,
			"message": "Event not found",
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Event retrieved successfully",
		"data":    event,
	})
}

func CreateEvent(context *gin.Context) {
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

	err = event.Store()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to create event",
			"data":    nil,
		})
		return
	}


	context.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "Event created successfully",
		"data":    event,
	})
}