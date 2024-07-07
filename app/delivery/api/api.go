package api

import (
	"forum/app/internal"
	"net/http"
)

func NewApiHandler(uc internal.Usecases) http.Handler {
	apiRouter := http.NewServeMux()

	// auth handles
	apiRouter.HandleFunc("POST /signin", nil)
	apiRouter.HandleFunc("POST /signup", nil)

	// user data handles
	apiRouter.HandleFunc("GET /profile/{profile_id}", nil)

	// post handles
	apiRouter.HandleFunc("POST /post", nil)
	apiRouter.HandleFunc("GET /post/{post_id}", nil)
	apiRouter.HandleFunc("POST /post/{post_id}/rate", nil)
	apiRouter.HandleFunc("GET /post/{post_id/comments}", nil)

	// comment handles
	apiRouter.HandleFunc("POST /post/{post_id}/comment", nil)
	apiRouter.HandleFunc("POST /post/{post_id}/comment/{comment_id}/rate", nil)

	return apiRouter
}
