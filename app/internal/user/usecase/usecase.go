package usecase

import "forum/app/internal/user"

type userUsecase struct {
	repo user.UserRepo
}

func NewUserService(repo user.UserRepo) user.UserService {
	return &userUsecase{repo: repo}
}
