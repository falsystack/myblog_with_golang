package usecases

import (
	"log"
	"toyproject_recruiting_community/dtos"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/repositories"
)

type PostUsecase interface {
	CreatePost(createPost dtos.CreatePost) error
}

type postUsecase struct {
	postRepository repositories.PostRepository
	userRepository repositories.UserRepository
}

func (p *postUsecase) CreatePost(createPost dtos.CreatePost) error {
	foundUser := p.userRepository.FindUserById(createPost.UserId)
	// Userフィールドはいるのか？
	postEntity := entities.Post{
		Title:   createPost.Title,
		Content: createPost.Content,
		User:    foundUser,
		UserId:  createPost.UserId,
	}
	err := p.postRepository.CreatePost(postEntity)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUsecase{postRepository: postRepository}
}
