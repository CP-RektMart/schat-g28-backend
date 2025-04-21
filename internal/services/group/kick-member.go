package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			kick group member
// @Tags			groups
// @Router			/api/v1/groups/{groupID}/members/{userID} [DELETE]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Param 			userID 		path 	uint 	true  "friend id"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleKickGroupMember(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.AddGroupMemberRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	group, err := h.repo.GetByID(req.GroupID, "Members")
	if err != nil {
		return errors.Wrap(err, "group not found")
	}

	if !group.IsOwner(userID) {
		return apperror.Forbidden("not an owner", nil)
	}

	if err := group.LeaveGroup(req.UserID); err != nil {
		return apperror.BadRequest(err.Error(), err)
	}

	if err := h.repo.LeaveGroup(group.ID, req.UserID); err != nil {
		return apperror.NotFound("user not found", err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
