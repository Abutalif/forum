package internal

import (
	"forum/app/internal/auth"
	"forum/app/internal/comment"
	"forum/app/internal/post"

	authRepo "forum/app/internal/auth/repo"
	commentRepo "forum/app/internal/comment/repo"
	postRepo "forum/app/internal/post/repo"
)

type Repos struct {
	auth.AuthRepo
	post.PostRepo
	comment.CommentRepo
}

func NewRepos() Repos {
	return Repos{
		AuthRepo:    authRepo.NewAuthRepo(),
		PostRepo:    postRepo.NewPostRepo(),
		CommentRepo: commentRepo.NewCommentRepo(),
	}
}
