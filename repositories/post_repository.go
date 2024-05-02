package repositories

import (
	"gorm.io/gorm"
	"log"
	"toyproject_recruiting_community/entities"
)

type PostRepository interface {
	CreatePost(post *entities.Post) error
	FindById(id uint) (*entities.Post, error)
	FindAll() (*[]entities.Post, error)
	RemoveById(id uint) error
	Update(updatePostEntity *entities.Post) (*entities.Post, error)
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

type postRepository struct {
	db *gorm.DB
}

func (pr *postRepository) Update(updatePostEntity *entities.Post) (*entities.Post, error) {
	tx := pr.db.Save(&updatePostEntity)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return updatePostEntity, nil
}

// TODO: Go言語のエラーハンドリングを調べてみる。もっと良い方法があるはず。
func (pr *postRepository) FindAll() (*[]entities.Post, error) {
	var posts []entities.Post
	tx := pr.db.Find(&posts)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return &posts, nil
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

func (pr *postRepository) CreatePost(post *entities.Post) error {
	result := pr.db.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
