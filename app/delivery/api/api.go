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

// TODO:
// 1) Implement path prefix
// 2) Implement method specification.

// then each struct serving as a handler (auth, post, etc.) will automatically
// decide what action to take depening on path starting with prefix
// and having a method
