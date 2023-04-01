package main

import (
	"file-storage/config"
	"file-storage/internal/app"
	"log"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	fileStorageApp := app.New()

	fileStorageApp.Run(cfg)
}
