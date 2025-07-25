package controllers

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/repositories"
	"example.com/rest-api/services"
	"example.com/rest-api/utils/responseformatter"
	"github.com/gin-gonic/gin"
)

var eventService = services.EventService{
	Repo: &repositories.EventRepository{},
}

func GetEvents(context *gin.Context) {
	events, err := eventService.GetEvents()
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to retrieve events")
		return
	}

	responseformatter.Success(context, http.StatusOK, "Events retrieved successfully", events)
}

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid event ID")
		return
	}

	event, err := eventService.GetEvent(id)
	if err != nil {
		responseformatter.Error(context, http.StatusNotFound, "Event not found")
		return
	}
	
	responseformatter.Success(context, http.StatusOK, "Event retrieved successfully", event)
}

func CreateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid input data")
		return
	}
	userId := context.GetInt64("userId")
	event.UserID = userId
	err = eventService.CreateEvent(&event)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to create event")
		return
	}

	responseformatter.Success(context, http.StatusCreated, "Event created successfully", event)
}

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid event ID")
		return
	}

	var event models.Event
	err = context.ShouldBindJSON(&event)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid input data")
		return
	}

	err = eventService.UpdateEvent(id, &event)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to update event")
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		responseformatter.Error(context, http.StatusForbidden, "You do not have permission to delete this event")
		return
	}

	responseformatter.Success(context, http.StatusOK, "Event updated successfully", event)
}

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid event ID")
		return
	}

	event, err := eventService.GetEvent(id)
	if err != nil {
		responseformatter.Error(context, http.StatusNotFound, "Event not found")
		return
	}

	userId := context.GetInt64("userId")
	if event.UserID != userId {
		responseformatter.Error(context, http.StatusForbidden, "You do not have permission to delete this event")
		return
	}

	err = eventService.DeleteEvent(event)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to delete event")
		return
	}

	responseformatter.Success(context, http.StatusOK, "Event deleted successfully", nil)
}