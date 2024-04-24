package usecases

import (
	"log"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/repositories/dtos"
)

type PostUsecase interface {
	CreatePost(createPost dtos.CreatePost) error
}

type postUsecase struct {
	postRepository repositories.PostRepository
	userRepository repositories.UserRepository
}

func (p *postUsecase) CreatePost(createPost dtos.CreatePost) error {
	err := p.postRepository.CreatePost(createPost)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUsecase{postRepository: postRepository}
}
