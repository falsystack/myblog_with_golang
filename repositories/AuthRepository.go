package repositories

import (
	"gorm.io/gorm"
	"toyproject_recruiting_community/entities"
)

type AuthRepository interface {
	FindById(id uint) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (ar *authRepository) FindById(id uint) (*entities.User, error) {
	return nil, nil
}
