package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			delete group
// @Tags			groups
// @Router			/api/v1/groups/{groupID} [DELETE]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeleteGroup(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.DeleteGroupRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	group, err := h.repo.GetByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed fetch group")
	}

	if !group.IsOwner(userID) {
		return apperror.Forbidden("user not an owner of the group", err)
	}

	if err := h.repo.Delete(req.ID); err != nil {
		return errors.Wrap(err, "failed delete group")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
