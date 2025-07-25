package controllers

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/repositories"
	"example.com/rest-api/services"
	"example.com/rest-api/utils"
	"example.com/rest-api/utils/responseformatter"
	"github.com/gin-gonic/gin"
)

var userService = services.UserService{
	Repo: &repositories.UserRepository{},
}

func StoreUser(context *gin.Context) {
	user := &models.User{}
	err := context.ShouldBindJSON(user)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid input data")
		return
	}

	err = userService.CreateUser(user)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to store user")
		return
	}

	responseformatter.Success(context, http.StatusCreated, "User created successfully", user)
}

func LoginUser(context *gin.Context) {
	user := &models.User{}
	err := context.ShouldBindJSON(user)
	if err != nil {
		responseformatter.Error(context, http.StatusBadRequest, "Invalid input data")
		return
	}

	err = userService.Authenticate(user)
	if err != nil {
		responseformatter.Error(context, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		responseformatter.Error(context, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	responseformatter.Success(context, http.StatusOK, "Login successful", gin.H{
		"token": token,
	})
}