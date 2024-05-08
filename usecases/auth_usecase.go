package usecases

import (
	"toyproject_recruiting_community/repositories"
	"toyproject_recruiting_community/usecases/output"
)

type AuthUsecase interface {
	FindByID(id uint) (*output.AuthResponse, error)
}

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{}
}

type authUsecase struct {
	authRepository repositories.AuthRepository
}

func (a *authUsecase) FindByID(id uint) (*output.AuthResponse, error) {
	user, err := a.authRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	return output.NewAuthResponse(
		user.ID,
		user.Name,
		user.Email,
	), nil
}
