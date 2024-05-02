package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	title   string `gorm:"not null"`
	content string `gorm:"not null"`
}

func (p *Post) GetTitle() string {
	return p.title
}

func (p *Post) GetContent() string {
	return p.content
}

func NewPost(title string, content string) *Post {
	return &Post{title: title, content: content}
}
