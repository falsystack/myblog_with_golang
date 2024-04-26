package usecases

import (
	"errors"
	"log"
	"toyproject_recruiting_community/repositories"
	rd "toyproject_recruiting_community/repositories/dtos"
	"toyproject_recruiting_community/response"
)

type PostUsecase interface {
	CreatePost(createPost rd.CreatePost) error
	FindById(id uint) (*response.PostResponse, error)
	RemoveById(id uint) error
	FindAll() (*[]response.PostResponse, error)
}

func NewPostUsecase(postRepository repositories.PostRepository) PostUsecase {
	return &postUsecase{postRepository: postRepository}
}

type postUsecase struct {
	postRepository repositories.PostRepository
	userRepository repositories.UserRepository
}

// TODO: pointerタイプとそうではないタイプの使い分けに関して調べてみる。
// 重いオブジェクトにだけポインタータイプを使うのか
func (pu *postUsecase) FindAll() (*[]response.PostResponse, error) {
	posts, err := pu.postRepository.FindAll()
	if err != nil {
		return nil, err
	}

	if len(*posts) < 1 {
		return nil, errors.New("no posts found")
	}
	var responses []response.PostResponse
	for _, post := range *posts {
		responses = append(responses, response.PostResponse{
			Title:   post.Title,
			Content: post.Content,
		})
	}
	return &responses, nil
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
