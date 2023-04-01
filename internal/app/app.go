package app

import (
	"file-storage/config"
	httpV1 "file-storage/internal/controller/http/v1"
	"file-storage/internal/domain/file_info"
	"file-storage/internal/storage/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

type FileStorageApp struct {
	httpServer *fiber.App
}

func New() *FileStorageApp {
	fiberServer := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	return &FileStorageApp{httpServer: fiberServer}
}

func (app *FileStorageApp) Run(cfg *config.Config) error {
	postgres.ConnectDb(cfg.PG.URL)
	defer postgres.CloseDb()

	fileInfoRepo := postgres.New(postgres.Database.Db)

	fileInfoService := file_info.New(fileInfoRepo)

	httpV1.NewRouter(app.httpServer, *fileInfoService)

	log.Println("Server running...")

	// Listen from a different goroutine
	go func() {
		if err := app.httpServer.Listen(":" + cfg.HTTP.Port); err != nil {
			log.Panic(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	_ = <-interrupt // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")

	shutdownTimeout := time.Duration(cfg.App.ShutdownTimeout) * time.Second
	_ = app.httpServer.ShutdownWithTimeout(shutdownTimeout)

	log.Println("Running cleanup tasks...")

	log.Println("Fiber was successful shutdown.")

	return nil
}
