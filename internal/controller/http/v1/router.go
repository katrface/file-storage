package v1

import (
	"file-storage/internal/domain/file_info"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(httpServer *fiber.App, service file_info.FileInfoService) {
	// K8s probe
	probHandler := func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	}
	httpServer.Get("/healthz", probHandler)

	// TODO Prometheus metrics https://github.com/ansrivas/fiberprometheus

	// Routers
	v1 := httpServer.Group("/api/v1")

	newFileInfoRoutes(v1, service)
}
