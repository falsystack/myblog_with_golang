package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toyproject_recruiting_community/controller/dtos"
	"toyproject_recruiting_community/usecases"
	ud "toyproject_recruiting_community/usecases/dtos"
)

type PostController interface {
	Create(ctx *gin.Context)
}

type postController struct {
	pu usecases.PostUsecase
}

func (pc *postController) Create(ctx *gin.Context) {
	var inputPost dtos.CreateInputPost
	if err := ctx.ShouldBind(&inputPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Go言語はJava(Spring)と違ってController側で開発者が直接Bindしているので
	// Front -> Controllerの間に使うDTOとController -> Usecaseの間に使うDTOの境界が曖昧だと感じた。
	createPost := ud.CreatePost{
		Title:   inputPost.Title,
		Content: inputPost.Content,
	}
	if err := pc.pu.CreatePost(createPost); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}

func NewPostController(pu usecases.PostUsecase) PostController {
	return &postController{pu: pu}
}
