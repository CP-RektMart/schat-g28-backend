package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			join group
// @Tags			groups
// @Router			/api/v1/groups/{groupID}/join [GET]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleJoinGroup(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.JoinGroupRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid-request", err)
	}

	group, err := h.repo.Get(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed fetching group")
	}

	if err := group.JoinGroup(userID); err != nil {
		return apperror.BadRequest(err.Error(), err)
	}

	if err := h.repo.JoinGroup(req.ID, userID); err != nil {
		return errors.Wrap(err, "failed saving membership")
	}

	return c.SendStatus(fiber.StatusOK)
}
