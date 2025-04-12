package file

import "github.com/gofiber/fiber/v2"

// @Summary			delete file
// @Tags			files
// @Router			/api/v1/files/{fileID} [DELETE]
// @Security		ApiKeyAuth
// @Param 			fileID 	path 	uint 	true "file id"
// @Success			200	
// @Failure			401	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeleteFile(c *fiber.Ctx) error {
	return nil
}
