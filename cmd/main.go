package main

import (
	"forum/app"
	"log"
	"os"
)

func main() {
	cfg := app.GetConfigs()
	logger := log.New(os.Stdout, "", log.LstdFlags)
	app := app.Init(cfg)

	if err := app.Run(logger); err != nil {
		logger.Println("Shutting Down due to error:", err)
	}
}
