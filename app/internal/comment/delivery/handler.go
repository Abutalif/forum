package delivery

import (
	"forum/app/internal/comment"
	"net/http"
)

type Handler struct {
	usecase comment.CommentService
}

func NewHandler(usecase comment.CommentService) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreateComment(rw http.ResponseWriter, r *http.Request) {}

func (h *Handler) RateComment(rw http.ResponseWriter, r *http.Request) {}
