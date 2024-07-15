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

func (h *Handler) CreatePost(rw http.ResponseWriter, r *http.Request) {}

func (h *Handler) RatePost(rw http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetPost(rw http.ResponseWriter, r *http.Request) {}

func (h *Handler) GetPostComments(rw http.ResponseWriter, r *http.Request) {}
