package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	server := gin.Default()

	server.GET("/ping", ping)

	server.Run(":8080")
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "pong",
		"data": nil,
	})
}