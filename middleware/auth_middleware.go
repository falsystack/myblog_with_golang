package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	"toyproject_recruiting_community/usecases"
)

func AuthMiddleware(au usecases.AuthUsecase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if len(header) == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
			ctx.Abort()
		}

		if !strings.HasPrefix(header, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
			ctx.Abort()
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			ctx.Abort()
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": jwt.ErrTokenExpired})
				ctx.Abort()
			}
			id := claims["sub"].(string)
			user, err := au.FindByID(id)
			if err != nil {
				// AbortなしでJSONだげでレスポンスを返すと{data: null}{error : "error message}変な形で返す。
				// ctx.JSON -> ctx.Abortで解決
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				ctx.Abort()
			}
			ctx.Set("user", user)
		}

		ctx.Next()

	}
}
