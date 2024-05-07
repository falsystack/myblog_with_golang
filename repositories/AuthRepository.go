package repositories

import "toyproject_recruiting_community/entities"

type AuthRepository interface {
	FindById(id uint) (*entities.User, error)
}
