package file

import (
	_ "github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			upload file
// @Tags			files
// @Router			/api/v1/files [POST]
// @Security		ApiKeyAuth
// @Param 			file 	formData 	file 	true  "file"
// @Success			200	{object}	dto.HttpResponse[dto.FileDetailResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUploadFile(c *fiber.Ctx) error {
	return nil
}
