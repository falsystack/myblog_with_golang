package repositories

import "toyproject_recruiting_community/entities"

type UserRepository interface {
	FindUserById(userId uint) entities.User
}
