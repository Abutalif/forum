package api

import (
	authHandler "forum/app/internal/auth/delivery"
	commentHandler "forum/app/internal/comment/delivery"
	postHandler "forum/app/internal/post/delivery"
	userHandler "forum/app/internal/user/delivery"

	"forum/app/internal"
	"net/http"
)

type Handlers struct {
	auth    *authHandler.Handler
	comment *commentHandler.Handler
	post    *postHandler.Handler
	user    *userHandler.Handler
}

func InitHandlers(uc internal.Usecases) *Handlers {
	return &Handlers{
		auth:    authHandler.NewHandler(uc.AuthService),
		comment: commentHandler.NewHandler(uc.CommentService),
		post:    postHandler.NewHandler(uc.PostService),
		user:    userHandler.NewHandler(uc.UserService),
	}
}

func (h *Handlers) NewApiHandler() http.Handler {
	apiRouter := http.NewServeMux()

	// auth handles
	apiRouter.HandleFunc("POST /signin", h.auth.SignIn)
	apiRouter.HandleFunc("POST /signup", h.auth.SignUp)

	// user data handles
	apiRouter.HandleFunc("GET /profile/{profile_id}", h.user.GetUserData)

	// TODO: add middleware for all bellow

	// post handles
	apiRouter.HandleFunc("POST /post", h.post.CreatePost)
	apiRouter.HandleFunc("GET /post/{post_id}", h.post.GetPost)
	apiRouter.HandleFunc("POST /post/{post_id}/rate", h.post.RatePost)
	apiRouter.HandleFunc("GET /post/{post_id/comments}", h.post.GetPostComments)

	// comment handles
	apiRouter.HandleFunc("POST /post/{post_id}/comment", h.comment.CreateComment)
	apiRouter.HandleFunc("POST /post/{post_id}/comment/{comment_id}/rate", h.comment.RateComment)

	return apiRouter
}
