package main

import (
	"forum/app"
	"log"
)

func main() {
	// I might also want to run a logger at the beggining.
	// Use Test-Driven development for this.
	cfg := app.GetConfigs()

	app := app.Init(cfg)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
