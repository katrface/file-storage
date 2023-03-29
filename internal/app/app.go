package app

import (
	"file-storage/config"
	httpV1 "file-storage/internal/controller/http/v1"
	"file-storage/internal/domain/file_info"
	"file-storage/internal/storage/postgres"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FileStorageApp struct {
	httpServer *fiber.App
}

func New() *FileStorageApp {
	fiberServer := fiber.New()

	return &FileStorageApp{httpServer: fiberServer}
}

func (app *FileStorageApp) Run(cfg *config.Config) error {
	postgres.ConnectDb(cfg.PG.URL)

	fileInfoRepo := postgres.New(postgres.Database.Db)

	fileInfoService := file_info.New(fileInfoRepo)

	httpV1.NewRouter(app.httpServer, *fileInfoService)

	fmt.Println("server running")

	app.httpServer.Listen(":" + cfg.HTTP.Port)

	return nil
}
