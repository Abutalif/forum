package api

import (
	"forum/app/internal"
	"net/http"
)

type ApiHandler struct {
	// here are what I have to handle

	// "POST /auth/signin"
	// "POST /auth/signup"

	// "GET /profile/{profile_id}"

	// "POST /post"
	// "GET /post/{post_id}"
	// "POST /post/{post_id}/like"
	// "POST /post/{post_id}/dislike"
	// "GET /post/{post_id}/comments"

	// "POST /post/{post_id}/comment"
	// "POST /post/{post_id}/comment/{comment_id}/like"
	// "POST /post/{post_id}/comment/{comment_id}/dislike"
}

func NewApiHandler(uc internal.Usecases) *ApiHandler {

	return &ApiHandler{
		// router.Handle("/auth", authHandler.NewHandler(uc.AuthService)),
		// router.Handle("/path", postHandler.NewHandler(uc.PostService)),
	}
}

func (a *ApiHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {}
