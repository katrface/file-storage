package v1

import (
	"errors"
	"file-storage/internal/domain/file_info"
	"file-storage/internal/pkg/api_errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func handleError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, file_info.ErrNotFound):
		apiError := api_errors.NewAPIError(file_info.ErrNotFound.Error(), api_errors.NotFoundErrorCode)

		return c.Status(http.StatusNotFound).JSON(apiError)

	case errors.Is(err, file_info.ErrRepository):
		apiError := api_errors.NewAPIError(file_info.ErrRepository.Error(), api_errors.DependencyErrorCode)

		return c.Status(http.StatusInternalServerError).JSON(apiError)

	default:
		apiError := api_errors.NewAPIError(err.Error(), api_errors.UnhandledErrorCode)

		return c.Status(http.StatusInternalServerError).JSON(apiError)
	}
}
