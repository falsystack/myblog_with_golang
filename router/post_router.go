package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"toyproject_recruiting_community/controller"
	"toyproject_recruiting_community/middleware"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases"
)

func PostRouter(r *gin.Engine, db *gorm.DB) {
	postRepository := repositories.NewPostRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepository)
	authUsecase := usecases.NewAuthUsecase(authRepository)
	postController := controller.NewPostController(postUsecase)

	postsRouter := r.Group("/posts", middleware.AuthMiddleware(authUsecase))
	postsRouter.POST("", postController.Create)
	postsRouter.GET("", postController.FindAll)
	postsRouter.GET("/:id", postController.FindById)
	postsRouter.PUT("/:id", postController.Update)
	postsRouter.DELETE("/:id", postController.RemoveById)
}
