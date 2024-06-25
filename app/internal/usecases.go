package internal

import (
	"forum/app/internal/auth"
	"forum/app/internal/post"
)

// create all usecases
type Usecases struct {
	auth.AuthService
	post.PostService
}

func NewUsecases(repos Repos) Usecases {
	return Usecases{}
}
