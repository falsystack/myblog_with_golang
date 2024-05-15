package usecases

import (
	"log"
	"toyproject_recruiting_community/entities"
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases/input"
	"toyproject_recruiting_community/usecases/output"
)

type AuthUsecase interface {
	FindByID(id string) (*output.AuthResponse, error)
	Create(createUser *input.CreateUser) error
}

func NewAuthUsecase(ar repositories.AuthRepository) AuthUsecase {
	return &authUsecase{ar: ar}
}

type authUsecase struct {
	ar repositories.AuthRepository
}

func (a *authUsecase) Create(createUser *input.CreateUser) error {
	err := a.ar.Create(entities.NewUser(createUser.ID, "", createUser.Email))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a *authUsecase) FindByID(id string) (*output.AuthResponse, error) {
	user, err := a.ar.FindById(id)
	if err != nil {
		return nil, err
	}

	return output.NewAuthResponse(
		user.ID,
		user.Name,
		user.Email,
	), nil
}
