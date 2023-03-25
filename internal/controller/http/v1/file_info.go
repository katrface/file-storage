package v1

import (
	"file-storage/internal/domain/file_info"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type fileInfoRoutes struct {
	service file_info.FileInfoService
}

func newFileInfoRoutes(router fiber.Router, service file_info.FileInfoService) {
	r := fileInfoRoutes{service: service}

	router.Post("/file-infos", r.createFileInfo)
	router.Get("/file-infos", r.getFileInfos)
	router.Get("/file-infos/:id", r.getFileInfoById)
	router.Delete("/file-infos/:id", r.removeFileInfoById)
}

func (r *fileInfoRoutes) createFileInfo(c *fiber.Ctx) error {
	var fileInfo file_info.FileInfo

	if err := c.BodyParser(&fileInfo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	createdFileInfo, _ := r.service.CreateFileInfo(fileInfo)
	return c.Status(http.StatusCreated).JSON(createdFileInfo)
}

func (r *fileInfoRoutes) getFileInfos(c *fiber.Ctx) error {
	fileInfos, _ := r.service.GetFileInfos()

	return c.Status(http.StatusOK).JSON(fileInfos)
}

func (r *fileInfoRoutes) getFileInfoById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	fileInfo, _ := r.service.GetFileInfoById(uint(id))

	return c.Status(http.StatusOK).JSON(fileInfo)
}

func (r *fileInfoRoutes) removeFileInfoById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	fileInfo, _ := r.service.RemoveFileInfoById(uint(id))

	return c.Status(http.StatusOK).JSON(fileInfo)
}
