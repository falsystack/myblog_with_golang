package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"toyproject_recruiting_community/controller"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/middleware"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases"
)

func main() {
	infra.Init()
	db := infra.SetupDB()
	r := gin.Default()

	postRouter(r, db)

	// auth
	authController := controller.NewAuthController()
	authRouter := r.Group("/auth")
	authRouter.GET("/google/login", authController.GoogleLogin)
	authRouter.GET("/google/callback", authController.GoogleAuthCallback)

	r.Run(":8080")

}

func postRouter(r *gin.Engine, db *gorm.DB) {
	postRepository := repositories.NewPostRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepository)
	authUsecase := usecases.NewAuthUsecase()
	postController := controller.NewPostController(postUsecase)

	postsRouter := r.Group("/posts", middleware.AuthMiddleware(authUsecase))
	postsRouter.POST("", postController.Create)
	postsRouter.GET("", postController.FindAll)
	postsRouter.GET("/:id", postController.FindById)
	postsRouter.PUT("/:id", postController.Update)
	postsRouter.DELETE("/:id", postController.RemoveById)
}
