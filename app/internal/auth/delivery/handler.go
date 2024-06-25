package delivery

import (
	"forum/app/internal/auth"
	"net/http"
)

type Handler struct {
	usecase auth.AuthService
}

func NewHandler(usecase auth.AuthService) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
