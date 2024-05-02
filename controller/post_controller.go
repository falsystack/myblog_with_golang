package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toyproject_recruiting_community/request"
	"toyproject_recruiting_community/usecases"
	"toyproject_recruiting_community/usecases/dtos/update"
	"toyproject_recruiting_community/usecases/input"
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
	inputPost := &input.Post{}
	if err := ctx.ShouldBind(inputPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.pu.Create(inputPost); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
