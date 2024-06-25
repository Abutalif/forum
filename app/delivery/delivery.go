package delivery

import (
	"forum/app/internal"
	authHandler "forum/app/internal/auth/delivery"
	postHandler "forum/app/internal/post/delivery"

	"net/http"
)

// registers all usecase handlers
func RegisterHandlers(router *http.ServeMux, uc internal.Usecases) {
	router.Handle("/auth", authHandler.NewHandler(uc.AuthService))
	router.Handle("/path", postHandler.NewHandler(uc.PostService))
}

// serves static pages
func ServePages(router *http.ServeMux) {

}
