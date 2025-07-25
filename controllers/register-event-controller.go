package controllers

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"example.com/rest-api/utils/responseformatter"
	"github.com/gin-gonic/gin"
)

func RegisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid event ID")
		return
	}

	event, err := eventService.GetEvent(eventId)
	if err != nil {
		responseformatter.Error(context, http.StatusNotFound, "Event not found")
		return
	}

	err = eventService.RegisterEvent(userId, event)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to register for event")
		return
	}

	responseformatter.Success(context, http.StatusCreated, "Successfully registered for event", nil)
}

func CancelEvent(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid event ID")
		return
	}
	event := &models.Event{}
	event.ID = eventId

	err = eventService.CheckUserEvent(userId, event)
	if err != nil {
		responseformatter.Error(context, http.StatusNotFound, "You are not registered for this event")
		return
	}

	err = eventService.CancelEvent(userId, event)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to cancel event registration")
		return
	}

	responseformatter.Success(context, http.StatusOK, "Successfully canceled event registration", nil)
}