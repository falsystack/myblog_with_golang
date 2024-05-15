package main

import (
	"github.com/gin-gonic/gin"
	"toyproject_recruiting_community/controller"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/router"
	"toyproject_recruiting_community/usecases"
)

func main() {
	infra.Init()
	db := infra.SetupDB()
	r := gin.Default()

	router.PostRouter(r, db)

	// auth
	repository := repositories.NewAuthRepository(db)
	usecase := usecases.NewAuthUsecase(repository)
	authController := controller.NewAuthController(usecase)

	authRouter := r.Group("/auth")
	authRouter.GET("/google/login", authController.GoogleLogin)
	authRouter.GET("/google/callback", authController.GoogleAuthCallback)

	r.Run(":8080")

}
