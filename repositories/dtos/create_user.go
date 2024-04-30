package dtos

import "toyproject_recruiting_community/entities"

// TODO: ファイル名変更
type CreateUser struct {
	Name     string
	Email    string
	Password string
	Posts    []entities.Post
}
