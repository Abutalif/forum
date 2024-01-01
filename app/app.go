package app

import (
	"fmt"
	"html/template"
	"net/http"
)

type app struct {
	data string
}

// then this will give the AppConfig struct to the app
// the app then will do "init(AppConfig)" (or "setup(AppConfig)", or "new(AppConfig)"), meaning the setup will be done ac
// then there will be app.Run() with error handling
// the "run()" shall also do graceful shutdown
// I should handle the error like:
// if err = app.Run(); err!=nil {
//
// {

func Init(cfg *config) *app {
	return &app{
		data: cfg.data1,
	}
}

func (a *app) Run() error {
	handler := http.NewServeMux()
	handler.HandleFunc("/", landingPage)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	server.ListenAndServe()
	return nil
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
