package main

import (
	"log"
	"rida-api/internal/config"
	"rida-api/internal/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := httpapp.NewApp(cfg)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
