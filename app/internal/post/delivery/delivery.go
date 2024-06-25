package delivery

import (
	"forum/app/internal/post"
	"net/http"
)

type Handler struct {
	usecase post.PostService
}

func NewHandler(usecase post.PostService) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {}
