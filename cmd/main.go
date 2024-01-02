package main

import (
	"forum/app"
	"log"
	"os"
)

// TODO: use tests
func main() {
	cfg := app.GetConfigs()
	logger := log.New(os.Stdout, "log", log.LstdFlags) // TODO: configure with cfg
	app := app.Init(cfg)

	if err := app.Run(logger); err != nil {
		logger.Fatal(err)
	}
}
