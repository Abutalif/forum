package delivery

import (
	"forum/app/internal/user"
	"net/http"
)

type Handler struct {
	usecase user.UserService
}

func NewHandler(usecase user.UserService) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetUserData(rw http.ResponseWriter, r *http.Request) {}
