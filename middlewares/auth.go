package middlewares

import (
	"net/http"
	"strings"

	"example.com/rest-api/utils"
	"example.com/rest-api/utils/responseformatter"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	{
		var userId int64
		var err error
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			responseformatter.MiddlewareError(context, http.StatusUnauthorized, "Authorization token is required")
			return
		}

		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			userId, err = utils.VerifyToken(token)
			if err != nil {
				responseformatter.MiddlewareError(context, http.StatusUnauthorized, err.Error())
				return
			}
		}

		context.Set("userId", userId)
		context.Next()
	}
}