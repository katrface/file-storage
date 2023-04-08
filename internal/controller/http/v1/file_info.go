package v1

import (
	"file-storage/internal/domain/file_info"
	"file-storage/internal/pkg/api_errors"
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
	router.Get("/file-infos/:id", r.getFileInfoByID)
	router.Delete("/file-infos/:id", r.removeFileInfoByID)
}

func (r *fileInfoRoutes) createFileInfo(c *fiber.Ctx) error {
	var fileInfo file_info.FileInfo

	if err := c.BodyParser(&fileInfo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	createdFileInfo, err := r.service.CreateFileInfo(fileInfo)
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(createdFileInfo)
}

func (r *fileInfoRoutes) getFileInfos(c *fiber.Ctx) error {
	fileInfos, err := r.service.GetFileInfos()
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fileInfos)
}

func (r *fileInfoRoutes) getFileInfoByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		appError := api_errors.NewAPIError("invalid file info id", api_errors.ValidationErrorCode)

		return c.Status(http.StatusUnprocessableEntity).JSON(appError)
	}

	fileInfo, err := r.service.GetFileInfoByID(uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fileInfo)
}

func (r *fileInfoRoutes) removeFileInfoByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		appError := api_errors.NewAPIError("invalid file info id", api_errors.ValidationErrorCode)

		return c.Status(http.StatusUnprocessableEntity).JSON(appError)
	}

	fileInfo, err := r.service.RemoveFileInfoByID(uint(id))
	if err != nil {
		return handleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(fileInfo)
}
