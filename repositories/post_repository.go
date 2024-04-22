package repositories

import (
	"toyproject_recruiting_community/entities"
)

type PostRepository interface {
	CreatePost(postEntity entities.Post) error
}
