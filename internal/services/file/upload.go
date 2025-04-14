package file

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
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
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return apperror.BadRequest("invalid file", err)
	}

	data, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "failed open file")
	}

	URL, err := h.store.UploadFile(
		ctx,
		"/files/"+file.Filename,
		file.Header.Get("content-type"),
		data,
		true,
	)
	if err != nil {
		return errors.Wrap(err, "failed upload file")
	}

	f, err := model.NewFile(URL, userID)
	if err != nil {
		return errors.Wrap(err, "failed create file")
	}

	f, err = h.repo.Create(f)
	if err != nil {
		return errors.Wrap(err, "failed save file record")
	}

	result := dto.ToFileDetailResponse(f)
	response := dto.HttpResponse[dto.FileDetailResponse]{
		Result: result,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
