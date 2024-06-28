package delivery

import (
	"forum/app/delivery/api"
	"forum/app/internal"

	"net/http"
)

// TODO: add choosing of API methods.

// FIXME: not sure if it is really good option. There must some interface
func RegisterApi(router *http.ServeMux, uc internal.Usecases) {
	router.Handle("/api/", http.StripPrefix("/api", api.NewApiHandler(uc)))
}

// serves static pages
func ServePages(router *http.ServeMux) {

}

// TODO: choose one of two:
// 1) delivery must have all handle, api has all HandlerFuncs.
// 2) api has all handles and HandlerFuncs. Delivery just registers high level api.
// 	But it also has control over which method to use: gin, gorilla, chi, stdlib.

// What is the easiest option?
// put every handlers in one place, register them where mux is defined.

// What is the hurdle?
// 	I don't like using nested ServeMux.
// 	I don't want to deal with switch-case and pattern recognition.
//  I also dislike the idea of using the HandlerFunc interface.
//	I am not sure if I want to or know how to properly use the Handler interface .

// BTW. If I am doing clean architecture, I can take inspiration from other's code,
// but on the section where database is used.
