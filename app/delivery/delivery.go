package delivery

import (
	"fmt"
	"forum/app/delivery/pages"
	"net/http"
)

func InitRoutes(mux *http.ServeMux) error {
	pages, err := pages.NewPages()
	if err != nil {
		return err
	}
	mux.Handle("/", pages)
	mux.HandleFunc("/api", placeholder)
	return nil
}

func placeholder(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("This does nothing")
}
