package auth

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Get me
// @Description		Get me
// @Tags			user
// @Router			/api/v1/me [GET]
// @Security		ApiKeyAuth
// @Success			200	{object}	dto.HttpResponse[dto.UserResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetMe(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	u, err := h.repo.GetUserByID(userID)
	if err != nil {
		return apperror.NotFound("user not found", err)
	}

	response := dto.HttpResponse[dto.UserResponse]{
		Result: dto.ToUserResponse(u),
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
