package delivery

import "net/http"

// registers all usecase handlers
func RegisterHandlers(router *http.ServeMux) {
	router.Handle("/auth", nil)
}

// serves static pages
func ServePages(router *http.ServeMux) {

}
