package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/", landingPage)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	server.ListenAndServe()
}

func landingPage(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("not get method")
		return
	}
	tpl, err := template.ParseGlob("./assets/templates/*.html")
	if err != nil {
		fmt.Println("parse glob: ", err)
		return
	}
	err = tpl.ExecuteTemplate(rw, "index.html", nil)
	if err != nil {
		fmt.Println("exec", err)
	}
}
