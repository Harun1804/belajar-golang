package controllers

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func StoreUser(context *gin.Context) {
	user := &models.User{}
	err := context.ShouldBindJSON(user)
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
	user := &models.User{}
	err := context.ShouldBindJSON(user)
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

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  false,
			"message": "Failed to generate token",
			"data":    nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Login successful",
		"data":    gin.H{
			"token": token,
		},
	})
}