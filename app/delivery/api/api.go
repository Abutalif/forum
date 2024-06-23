package api

import (
	"fmt"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.Handle("/api", &auth{})
	mux.Handle("/api/post", &post{})
}

type auth struct{}

func (a *auth) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("do smth")
}

type post struct{}

func (p *post) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("do smth")
}
