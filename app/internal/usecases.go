package internal

import (
	"forum/app/internal/auth"
	"forum/app/internal/comment"
	"forum/app/internal/post"
	"forum/app/internal/user"

	authUC "forum/app/internal/auth/usecase"
	commentUC "forum/app/internal/comment/usecase"
	postUC "forum/app/internal/post/usecase"
	userUC "forum/app/internal/user/usecase"
)

// TODO: separate UC by name
// create all usecases
type Usecases struct {
	auth.AuthService
	post.PostService
	comment.CommentService
	user.UserService
}

func NewUsecases(repos Repos) Usecases {
	return Usecases{
		AuthService:    authUC.NewAuthService(repos.AuthRepo),
		PostService:    postUC.NewPostService(repos.PostRepo),
		CommentService: commentUC.NewPostService(repos.CommentRepo),
		UserService:    userUC.NewUserService(repos.UserRepo),
	}
}
