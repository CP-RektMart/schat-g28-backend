package friend

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// @Summary			list user's friends
// @Tags			friends
// @Router			/api/v1/friends [GET]
// @Security		ApiKeyAuth
// @Success			200 {object}	dto.HttpListResponse[dto.UserResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListFriends(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	user, err := h.authRepo.GetUserByID(userID, "Friends")
	if err != nil {
		return errors.Wrap(err, "failed fetch user")
	}

	result := dto.ToUsersReponse(user.Friends)

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.UserResponse]{
		Result: result,
	})
}
