package internal

import (
	"forum/app/internal/auth"
	"forum/app/internal/comment"
	"forum/app/internal/post"
	"forum/app/internal/user"

	authRepo "forum/app/internal/auth/repo"
	commentRepo "forum/app/internal/comment/repo"
	postRepo "forum/app/internal/post/repo"
	userRepo "forum/app/internal/user/repo"
)

type Repos struct {
	auth.AuthRepo
	post.PostRepo
	comment.CommentRepo
	user.UserRepo
}

func NewRepos() Repos {
	return Repos{
		AuthRepo:    authRepo.NewAuthRepo(),
		PostRepo:    postRepo.NewPostRepo(),
		CommentRepo: commentRepo.NewCommentRepo(),
		UserRepo:    userRepo.NewUserRepo(),
	}
}
