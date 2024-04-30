package repositories

import (
	"gorm.io/gorm"
	"log"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/repositories/dtos"
)

// TODO: ファイル名変更
type UserRepository interface {
	Create(createUser dtos.CreateUser) error
	FindUserById(userId uint) entities.User
}

type userRepository struct {
	db *gorm.DB
}

func (ur *userRepository) Create(createUser dtos.CreateUser) error {
	user := entities.User{
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: createUser.Password,
	}
	tx := ur.db.Create(&user)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
	}
	return nil
}

func (ur *userRepository) FindUserById(userId uint) entities.User {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
