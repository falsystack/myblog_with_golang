package repositories

import (
	"gorm.io/gorm"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/repositories/dtos"
)

type PostRepository interface {
	CreatePost(createPost dtos.CreatePost) error
	FindById(id uint) (*entities.Post, error)
	RemoveById(id uint) error
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

type postRepository struct {
	db *gorm.DB
}

func (pr *postRepository) RemoveById(id uint) error {
	post, err := pr.FindById(id)
	if err != nil {
		return err
	}

	tx := pr.db.Unscoped().Delete(post)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (pr *postRepository) FindById(id uint) (*entities.Post, error) {
	var post entities.Post
	tx := pr.db.Find(&post, "id = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &post, nil
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
