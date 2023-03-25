package main

import (
	"file-storage/internal/app"
)

func main() {
	fileStorageApp := app.New()

	fileStorageApp.Run()
}
