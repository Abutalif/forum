package internal

import (
	"forum/app/internal/auth"
	authUC "forum/app/internal/auth/usecase"
	"forum/app/internal/post"
	postUC "forum/app/internal/post/usecase"
)

// create all usecases
type Usecases struct {
	auth.AuthService
	post.PostService
}

func NewUsecases(repos Repos) Usecases {
	return Usecases{
		AuthService: authUC.NewAuthService(repos.AuthRepo),
		PostService: postUC.NewPostService(repos.PostRepo),
	}
}
