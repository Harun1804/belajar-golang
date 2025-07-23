package controllers

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func StoreUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid input data",
			"data":    nil,
		})
		return
	}

	err = user.CreateUser()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to store user",
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status":  true,
		"message": "User created successfully",
		"data":    user,
	})
}

func LoginUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid input data",
			"data":    nil,
		})
		return
	}

	err = user.Authenticate()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login successful",
		"data":    nil,
	})
}