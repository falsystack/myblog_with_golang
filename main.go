package main

import (
	"github.com/gin-gonic/gin"
	"toyproject_recruiting_community/controller"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases"
)

func main() {
	infra.Init()
	db := infra.SetupDB()

	repository := repositories.NewPostRepository(db)
	usecase := usecases.NewPostUsecase(repository)
	postController := controller.NewPostController(usecase)

	r := gin.Default()
	postsRouter := r.Group("/posts")
	postsRouter.POST("", postController.Create)

	r.Run(":8080")

}
