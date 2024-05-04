package usecases

import (
	"log"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases/input"
	"toyproject_recruiting_community/usecases/output"
)

type PostUsecase interface {
	Create(inputPost *input.CreatePost) error
	FindById(id uint) (*output.PostResponse, error)
	Update(updatePost *input.UpdatePost) (*output.PostResponse, error)
	FindAll() ([]output.PostResponse, error)
	RemoveById(id uint) error
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUsecase{postRepository: postRepository}
}

type postUsecase struct {
	postRepository repositories.PostRepository
}

func (pu *postUsecase) Update(updatePost *input.UpdatePost) (*output.PostResponse, error) {
	foundPost, err := pu.postRepository.FindById(updatePost.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	foundPost.Update(updatePost)

	updatedPost, err := pu.postRepository.Update(foundPost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &output.PostResponse{
		ID:      updatedPost.ID,
		Title:   updatedPost.Title,
		Content: updatedPost.Content,
	}, nil
}

func (pu *postUsecase) FindAll() ([]output.PostResponse, error) {
	posts, err := pu.postRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []output.PostResponse
	for _, post := range *posts {
		responses = append(responses, output.PostResponse{
			ID:      post.ID,
			Title:   post.Title,
			Content: post.Content,
		})
	}

	return responses, nil
}

func (pu *postUsecase) RemoveById(id uint) error {
	return pu.postRepository.RemoveById(id)
}

func (pu *postUsecase) FindById(id uint) (*output.PostResponse, error) {
	foundPost, err := pu.postRepository.FindById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp := &output.PostResponse{
		ID:      foundPost.ID,
		Title:   foundPost.Title,
		Content: foundPost.Content,
	}

	return resp, nil
}

func (pu *postUsecase) Create(inputPost *input.CreatePost) error {
	post := entities.NewPost(inputPost.Title, inputPost.Content)
	err := pu.postRepository.CreatePost(post)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
