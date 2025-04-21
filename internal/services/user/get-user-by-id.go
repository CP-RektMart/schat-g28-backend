package user

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			get user by id
// @Tags			user
// @Router			/api/v1/users/{id} [GET]
// @Security		ApiKeyAuth
// @Param 			id 	path 	uint 	true  "user id"
// @Success			200 {object} 	dto.HttpResponse[dto.UserDetailResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetByID(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get userID from context")
	}

	var req dto.GetUserReqest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	user, err := h.repo.GetByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed fetch user")
	}

	var messages []model.DirectMessage

	if err := h.db.
		Order("created_at").
		Joins("Receiver").
		Joins("Sender").
		Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", userID, req.ID, req.ID, userID).
		Find(&messages).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.NotFound("Messages not found", err)
		}
		return errors.Wrap(err, "failed to get sent messages")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":     user,
		"messages": messages,
	})
}
