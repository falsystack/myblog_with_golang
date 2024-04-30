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

// TODO: フィいる名をGoらしく変える。
// TODO: もっと良いメソッド名を考える。
type PostController interface {
	Create(ctx *gin.Context)
	// TODO: IDも一緒に返却するように修正
	FindPostById(ctx *gin.Context)
	// TODO: IDも一緒に返却するように修正
	FindAllPosts(ctx *gin.Context)
	Remove(ctx *gin.Context)
	// TODO: idはpath variableとして受け取るように修正
	Update(ctx *gin.Context)
}

func NewPostController(pu usecases.PostUsecase) PostController {
	return &postController{pu: pu}
}

type postController struct {
	pu usecases.PostUsecase
}

func (pc *postController) Update(ctx *gin.Context) {
	var requestPost request.UpdatePost
	if err := ctx.ShouldBindJSON(&requestPost); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, _ := strconv.ParseUint(string(requestPost.ID), 10, 64)
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

func (pc *postController) FindAllPosts(ctx *gin.Context) {
	posts, err := pc.pu.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": posts})
}

func (pc *postController) Remove(ctx *gin.Context) {
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

func (pc *postController) FindPostById(ctx *gin.Context) {
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
	if err := pc.pu.CreatePost(createPost); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusCreated)
}
