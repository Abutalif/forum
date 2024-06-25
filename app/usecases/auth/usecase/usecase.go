package usecase

import (
	"forum/app/entities"
	"forum/app/usecases/auth"
)

type authUsecase struct {
	// logger *log.Logger
	repo auth.AuthRepo
}

func NewAuthService(repo auth.AuthRepo) auth.AuthService {
	return &authUsecase{
		repo: repo,
	}
}

func (a *authUsecase) SignIn(user entities.User) error {
	return nil
}

func (a *authUsecase) SignUp() error {
	return nil
}
