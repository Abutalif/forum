package internal

import (
	"forum/app/internal/auth"
	"forum/app/internal/post"

	authRepo "forum/app/internal/auth/repo"
	postRepo "forum/app/internal/post/repo"
)

type Repos struct {
	auth.AuthRepo
	post.PostRepo
}

func NewRepos() Repos {
	return Repos{
		AuthRepo: authRepo.NewAuthRepo(),
		PostRepo: postRepo.NewPostRepo(),
	}
}
