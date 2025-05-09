package auth

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Logout
// @Tags			auth
// @Router			/api/v1/auth/logout [POST]
// @Security		ApiKeyAuth
// @Success			200
// @Failure			401	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleLogout(c *fiber.Ctx) error {
	ctx := c.UserContext()

	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}

	if err := h.jwtService.RemoveToken(ctx, userID); err != nil {
		return errors.Wrap(err, "failed to remove token")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
