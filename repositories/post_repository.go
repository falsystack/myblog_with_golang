package repositories

import (
	"gorm.io/gorm"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/usecases/dtos"
)

type PostRepository interface {
	CreatePost(createPost dtos.CreatePost) error
}

type postRepository struct {
	db *gorm.DB
}

func (pr *postRepository) CreatePost(createPost dtos.CreatePost) error {
	postEntity := entities.Post{
		Title:   createPost.Title,
		Content: createPost.Content,
	}
	result := pr.db.Create(&postEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}
