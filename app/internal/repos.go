package internal

import (
	"forum/app/internal/auth"
	"forum/app/internal/post"
)

type Repos struct {
	auth.AuthRepo
	post.PostRepo
}

func NewRepos() Repos {
	return Repos{}
}
