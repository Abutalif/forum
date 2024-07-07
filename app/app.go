package app

import (
	"context"
	"forum/app/delivery"
	"forum/app/delivery/api"
	"forum/app/internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type app struct {
	serveStatic bool
	address     string
}

func Init(cfg *config) *app {
	// TODO: init must be a bit more useful
	return &app{
		serveStatic: cfg.serveStatic,
		address:     cfg.address + ":" + cfg.port,
	}
}

func (a *app) Run(logger *log.Logger) error {
	repos := internal.NewRepos()
	usecases := internal.NewUsecases(repos)

	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", api.NewApiHandler(usecases)))

	if a.serveStatic {
		delivery.ServePages(router)
	}

	server := &http.Server{
		Addr:    a.address,
		Handler: router,
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
