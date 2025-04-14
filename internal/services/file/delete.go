package file

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

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
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.DeleteFileRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid-request", err)
	}

	if err := h.repo.Delete(ctx, req.ID, func(f model.File) error {
		return f.IsOwner(userID)
	}); err != nil {
		return errors.Wrap(err, "failed delete ")
	}

	return nil
}
