package friend

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

// @Summary			unfriend
// @Tags			friends
// @Router			/api/v1/friends/{friendID} [DELETE]
// @Security		ApiKeyAuth
// @Param 			friendID 	path 	uint 	true  "friend id"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUnFriend(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.UnFriendRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	user, err := h.authRepo.GetUserByID(userID, "Friends")
	if err != nil {
		return errors.Wrap(err, "failed fetch user")
	}

	if err := user.CanUnFriend(req.FriendID); err != nil {
		return apperror.BadRequest(err.Error(), err)
	}

	if err := h.authRepo.UnFriend(userID, req.FriendID); err != nil {
		return errors.Wrap(err, "failed update storage")
	}

	return c.SendStatus(fiber.StatusOK)
}
