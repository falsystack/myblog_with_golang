package main

import (
	"github.com/gin-gonic/gin"
	"toyproject_recruiting_community/controller"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/router"
)

func main() {
	infra.Init()
	db := infra.SetupDB()
	r := gin.Default()

	router.PostRouter(r, db)

	// auth
	authController := controller.NewAuthController()
	authRouter := r.Group("/auth")
	authRouter.GET("/google/login", authController.GoogleLogin)
	authRouter.GET("/google/callback", authController.GoogleAuthCallback)

	r.Run(":8080")

}
