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
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if len(header) == 0 {
			log.Println("No Authorization header")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		if !strings.HasPrefix(header, "Bearer ") {
			log.Println("Invalid Authorization header")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("Unexpected signing method")
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": jwt.ErrTokenExpired})
				return
			}
		}

		// TODO: Find User

		// TODO: Set User
		//ctx.Set("user", user)
		ctx.Next()

	}
}
