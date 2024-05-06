package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exists := ctx.Get("user")
		if !exists {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
