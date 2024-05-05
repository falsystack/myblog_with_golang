package main

import (
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
	"toyproject_recruiting_community/usecases/output"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatal("Error loading .env.test file")
	}

	os.Exit(m.Run())
}

func setupTestData(db *gorm.DB) {
	posts := []*entities.Post{
		{Title: "テストアイテム1", Content: "", UserID: 1},
		{Title: "テストアイテム2", Content: "テスト2", UserID: 1},
		{Title: "テストアイテム3", Content: "テスト3", UserID: 2},
	}

	users := []*entities.User{
		{
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

func setup() *gin.Engine {
	db := infra.SetupDB()
	db.AutoMigrate(&entities.User{}, &entities.Post{})

	setupTestData(db)
	r := gin.Default()

	postRouter(r, db)
	return r
}

func TestPostController(t *testing.T) {
	// given
	r := setup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/posts", nil)

	// when
	r.ServeHTTP(w, req)

	var resp map[string][]output.PostResponse
	json.Unmarshal([]byte(w.Body.String()), &resp)

	fmt.Println(w.Body)
	// then
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 3, len(resp["data"]))
}
