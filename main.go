package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"toyproject_recruiting_community/controller"
	_ "toyproject_recruiting_community/docs"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/router"
	"toyproject_recruiting_community/usecases"
)

// @title Toy Project Recruiting Community
// @version 0.10
// @description このswaggerはgin-swaggerにより生成されました。
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")

}
