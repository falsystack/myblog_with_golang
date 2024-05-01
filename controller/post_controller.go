package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toyproject_recruiting_community/controller/dtos"
	ud "toyproject_recruiting_community/repositories/dtos"
	"toyproject_recruiting_community/request"
	"toyproject_recruiting_community/usecases"
	"toyproject_recruiting_community/usecases/dtos/update"
)

/**
悩み
- Controller の宣言部のみをみると何を返しいるかが把握できない
- 注釈をつけて（つけたくないが）何を返しているかを見えるようにする
*/

type PostController interface {
	Create(ctx *gin.Context)

	// FindById return response.PostResponse
	FindById(ctx *gin.Context)

	// FindAll return []response.PostResponse
	FindAll(ctx *gin.Context)

	// Update return response.PostResponse
	Update(ctx *gin.Context)

	RemoveById(ctx *gin.Context)
}

func NewPostController(pu usecases.PostUsecase) PostController {
	return &postController{pu: pu}
}

type postController struct {
	pu usecases.PostUsecase
}

func (pc *postController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var requestPost request.UpdatePost
	if err := ctx.ShouldBindJSON(&requestPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatePost := update.UpdatePost{
		ID:      uint(id),
		Title:   requestPost.Title,
		Content: requestPost.Content,
	}
	resp, err := pc.pu.Update(updatePost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (pc *postController) FindAll(ctx *gin.Context) {
	posts, err := pc.pu.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func (pc *postController) RemoveById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Post Id"})
		return
	}

	err = pc.pu.RemoveById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func (pc *postController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Post Id"})
		return
	}

	postResponse, err := pc.pu.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": postResponse})
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
	if err := pc.pu.Create(createPost); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
