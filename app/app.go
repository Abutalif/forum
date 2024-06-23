package app

import (
	"context"
	"forum/app/delivery"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	delivery.InitRoutes(handler)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Println("Shutting down due to:", err)
		}
	}()

	// gracefull shutdown
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	exitSig := <-exit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Println("Recieved terminate, graceful shutdown. Signal:", exitSig)
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalln("Could not gracefully shutdown:", err)
	}
	return nil
}
