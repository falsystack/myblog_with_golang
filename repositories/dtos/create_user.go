package dtos

import "toyproject_recruiting_community/entities"

type CreateUser struct {
	Name     string
	Email    string
	Password string
	Posts    []entities.Post
}
