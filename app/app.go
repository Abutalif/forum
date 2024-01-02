package app

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type app struct {
	data string
}

func Init(cfg *config) *app {
	return &app{
		data: cfg.data1,
	}
}

func (a *app) Run(logger *log.Logger) error {
	handler := http.NewServeMux()
	handler.HandleFunc("/", landingPage)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
		// TODO: add other fields if necessary
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	signal.Notify(exit, os.Kill)
	exitSig := <-exit
	ctx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	logger.Println("Recieved terminate, graceful shutdown", exitSig)
	server.Shutdown(ctx)
	defer shutdown()
	return nil
}

// TODO: add handlers
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
