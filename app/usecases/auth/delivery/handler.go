package delivery

import (
	"forum/app/usecases/auth"
	"net/http"
)

type AuthHandler struct {
	usecase auth.AuthService
}

func (a *AuthHandler) SignIn(rw http.ResponseWriter, r *http.Request) {

}
