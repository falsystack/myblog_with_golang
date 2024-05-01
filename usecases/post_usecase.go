package usecases

import (
	"errors"
	"log"
	"toyproject_recruiting_community/repositories"
	rd "toyproject_recruiting_community/repositories/dtos"
	"toyproject_recruiting_community/response"
	"toyproject_recruiting_community/usecases/dtos/update"
)

type PostUsecase interface {
	Create(createPost rd.CreatePost) error
	FindById(id uint) (*response.PostResponse, error)
	FindAll() ([]*response.PostResponse, error)
	Update(updatePost update.UpdatePost) (*response.PostResponse, error)
	RemoveById(id uint) error
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUsecase{postRepository: postRepository}
}

type postUsecase struct {
	postRepository repositories.PostRepository
	userRepository repositories.UserRepository
}

func (pu *postUsecase) Update(updatePost update.UpdatePost) (*response.PostResponse, error) {
	foundPost, err := pu.postRepository.FindById(updatePost.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	foundPost.Title = updatePost.Title
	foundPost.Content = updatePost.Content

	updatedPost, err := pu.postRepository.Update(foundPost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response.NewPostResponse(
		updatePost.ID,
		updatePost.Title,
		updatedPost.Content,
	), nil
}

func (pu *postUsecase) FindAll() ([]*response.PostResponse, error) {
	posts, err := pu.postRepository.FindAll()
	if err != nil {
		return nil, err
	}

	if len(*posts) < 1 {
		return nil, errors.New("Post Not Found")
	}

	var responses []*response.PostResponse
	for _, post := range *posts {
		responses = append(responses, response.NewPostResponse(
			post.ID,
			post.Title,
			post.Content,
		))
	}
	return responses, nil
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

	return response.NewPostResponse(
		foundPost.ID,
		foundPost.Title,
		foundPost.Content), nil
}

func (pu *postUsecase) Create(createPost rd.CreatePost) error {
	err := pu.postRepository.CreatePost(createPost)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
