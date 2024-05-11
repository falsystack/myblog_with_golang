package entities

import (
	"gorm.io/gorm"
	"toyproject_recruiting_community/entities/utils"
	"toyproject_recruiting_community/usecases/input"
)

type Post struct {
	gorm.Model
	Title   string `json:"title" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	UserID  string `json:"user_id" gorm:"not null"`
}

func NewPost(title string, content string) *Post {
	return &Post{Title: title, Content: content}
}

func (p *Post) Update(updatePost *input.UpdatePost) {
	if updatePost == nil {
		return
	}

	if !utils.IsEmptyString(updatePost.Title) {
		p.Title = updatePost.Title
	}
	if !utils.IsEmptyString(updatePost.Content) {
		p.Content = updatePost.Content
	}
}
