package usecases

import (
	"log"
	"toyproject_recruiting_community/repositories"
	rd "toyproject_recruiting_community/repositories/dtos"
	"toyproject_recruiting_community/response"
)

type PostUsecase interface {
	CreatePost(createPost rd.CreatePost) error
	FindById(id uint) (*response.PostResponse, error)
	RemoveById(id uint) error
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUsecase{postRepository: postRepository}
}

type postUsecase struct {
	postRepository repositories.PostRepository
	userRepository repositories.UserRepository
}

func (pu *postUsecase) RemoveById(id uint) error {
	return pu.postRepository.RemoveById(id)
}

func (pu *postUsecase) FindById(id uint) (*response.PostResponse, error) {
	foundPost, err := pu.postRepository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &response.PostResponse{
		Title:   foundPost.Title,
		Content: foundPost.Content,
	}, nil
}

func (pu *postUsecase) CreatePost(createPost rd.CreatePost) error {
	err := pu.postRepository.CreatePost(createPost)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
