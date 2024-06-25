package repo

import (
	"forum/app/entities"
	"forum/app/usecases/auth"
)

type authRepo struct {
}

func NewAuthRepo() auth.AuthRepo {
	return &authRepo{}
}

func (a *authRepo) GetUser(name string) entities.User {
	return entities.User{}
}
