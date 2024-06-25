package auth

import (
	"forum/app/entities"
)

type AuthService interface {
	SignIn(entities.User) error
	SignUp() error
}

type AuthRepo interface {
	GetUser(username string) entities.User
}
