package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toyproject_recruiting_community/usecases"
	"toyproject_recruiting_community/usecases/input"
)

/**
悩み
- Controller の宣言部のみをみると何を返しいるかが把握できない
- 注釈をつけて（つけたくないが）何を返しているかを見えるようにする
*/

type PostController interface {
	Create(ctx *gin.Context)
	FindById(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	RemoveById(ctx *gin.Context)
}

func NewPostController(pu usecases.PostUsecase) PostController {
	return &postController{pu: pu}
}

type postController struct {
	pu usecases.PostUsecase
}

// Update return output.PostResponse
func (pc *postController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updatePost input.UpdatePost
	updatePost.ID = uint(id)
	if err := ctx.ShouldBindJSON(&updatePost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := pc.pu.Update(&updatePost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// FindAll return []output.PostResponse
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CreatePost Id"})
		return
	}

	err = pc.pu.RemoveById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

// FindById return output.PostResponse
func (pc *postController) FindById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid CreatePost Id"})
		return
	}

	postResponse, err := pc.pu.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": postResponse})
}

// Create receives input.CreatePost as an argument and generate entities.Post
func (pc *postController) Create(ctx *gin.Context) {
	createPost := &input.CreatePost{}
	if err := ctx.ShouldBind(createPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.pu.Create(createPost); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
