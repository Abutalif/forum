package repo

import "forum/app/internal/user"

type userRepo struct{}

func NewUserRepo() user.UserRepo {
	return &userRepo{}
}
