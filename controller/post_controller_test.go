package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/infra"
	"toyproject_recruiting_community/middleware"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases"
	"toyproject_recruiting_community/usecases/input"
	"toyproject_recruiting_community/usecases/output"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}

	os.Exit(m.Run())
}

func setupTestData(db *gorm.DB) {
	posts := []*entities.Post{
		{Title: "テストアイテム1", Content: "", UserID: "94A803A5-82BA-4BBB-B597-DE97569A4F3C"},
		{Title: "テストアイテム2", Content: "テスト2", UserID: "94A803A5-82BA-4BBB-B597-DE97569A4F3C"},
		{Title: "テストアイテム3", Content: "テスト3", UserID: "94A803A5-82BA-4BBB-B597-DE97569A4F3C"},
	}

	users := []*entities.User{
		{
			ID:    "94A803A5-82BA-4BBB-B597-DE97569A4F3C",
			Name:  "testuser1",
			Email: "test1@example.com",
		},
		{
			Name:  "testuser2",
			Email: "test2@example.com",
		},
	}

	for _, user := range users {
		fmt.Println(user)
		db.Create(&user)
	}
	for _, item := range posts {
		db.Create(&item)
	}
}

func setup() *gorm.DB {
	db := infra.SetupDB()
	db.Migrator().DropTable(&entities.User{}, &entities.Post{})
	db.AutoMigrate(&entities.User{}, &entities.Post{})

	setupTestData(db)
	//r := gin.Default()

	//return r
	return db
}

func TestCreatePost(t *testing.T) {
	// given
	db := setup()
	postRepository := repositories.NewPostRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepository)
	postController := NewPostController(postUsecase)

	createPost := input.CreatePost{
		Title:   "Test Post",
		Content: "Test Content",
	}
	reqBody, _ := json.Marshal(createPost)

	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	_, router := gin.CreateTestContext(w)

	// when
	router.POST("/posts", postController.Create)
	router.ServeHTTP(w, req)

	// then
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestFindByID(t *testing.T) {
	// given
	db := setup()
	postRepository := repositories.NewPostRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepository)
	postController := NewPostController(postUsecase)

	req := httptest.NewRequest("GET", "/posts/3", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	_, r := gin.CreateTestContext(w)

	// when
	r.GET("/posts/:id", postController.FindById)
	r.ServeHTTP(w, req)

	// then
	var res map[string]output.PostResponse
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "テストアイテム3", res["data"].Title)
	assert.Equal(t, uint(3), res["data"].ID)
}

func TestFindAll(t *testing.T) {
	// given
	db := setup()
	postRepository := repositories.NewPostRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepository)
	postController := NewPostController(postUsecase)

	req := httptest.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()

	_, r := gin.CreateTestContext(w)

	// when
	r.GET("/posts", postController.FindAll)
	r.ServeHTTP(w, req)

	// then
	var res map[string][]output.PostResponse
	json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(res["data"]))
	assert.Equal(t, res["data"][0].Title, "テストアイテム1")
}

func TestUpdate(t *testing.T) {
	// given
	db := setup()
	postRepository := repositories.NewPostRepository(db)
	authRepository := repositories.NewAuthRepository(db)
	postUsecase := usecases.NewPostUsecase(postRepository)
	authUsecase := usecases.NewAuthUsecase(authRepository)
	postController := NewPostController(postUsecase)
	authMiddleware := middleware.AuthMiddleware(authUsecase)

	updatePost := input.UpdatePost{
		ID:      3,
		Title:   "テストアイテム1",
		Content: "テスト1",
	}
	reqBody, _ := json.Marshal(updatePost)

	req := httptest.NewRequest("PUT", "/posts/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	_, r := gin.CreateTestContext(w)

	// when
	r.PUT("/posts/:id", authMiddleware, postController.Update)
	r.ServeHTTP(w, req)

	// then
	var res map[string]output.PostResponse
	json.Unmarshal([]byte(w.Body.String()), &res)

	fmt.Println(w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "テスト1", res["data"].Content)
}
