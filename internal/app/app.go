package app

import (
	httpV1 "file-storage/internal/controller/http/v1"
	"file-storage/internal/domain/file_info"
	"file-storage/internal/storage"
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

func (app *FileStorageApp) Run() error {
	storage.ConnectDb()

	fileInfoRepo := storage.New(storage.Database.Db)

	fileInfoService := file_info.New(fileInfoRepo)

	httpV1.NewRouter(app.httpServer, *fileInfoService)

	fmt.Println("server running")

	app.httpServer.Listen(":3000")

	return nil
}
