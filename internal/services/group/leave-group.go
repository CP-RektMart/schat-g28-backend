package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			leave group
// @Tags			groups
// @Router			/api/v1/groups/{groupID}/leave [GET]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @success			204 
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleLeaveGroup(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.LeaveGroupRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid-request", err)
	}

	group, err := h.repo.Get(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed fetching group")
	}

	if err := group.LeaveGroup(userID); err != nil {
		return apperror.BadRequest(err.Error(), err)
	}

	if err := h.repo.LeaveGroup(req.ID, userID); err != nil {
		return errors.Wrap(err, "failed saving membership")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
