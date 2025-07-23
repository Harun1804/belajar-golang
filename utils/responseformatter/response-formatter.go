package responseformatter

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
    Status  bool        `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func Success(context *gin.Context, statusCode int, message string, data interface{}) {
    resp := Response{
        Status:  true,
        Message: message,
        Data:    data,
    }
    context.JSON(statusCode, resp)
}

func Error(context *gin.Context, statusCode int, message string) {
    resp := Response{
        Status:  false,
        Message: message,
        Data:    nil,
    }
    context.JSON(statusCode, resp)
}

func ValidationError(context *gin.Context, statusCode int, message string, errors map[string]string) {
    resp := Response{
        Status:  false,
        Message: message,
        Data:    errors,
    }
    context.JSON(statusCode, resp)
}