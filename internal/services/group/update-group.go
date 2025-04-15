package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			update group
// @Tags			groups
// @Router			/api/v1/groups/{groupID} [PATCH]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Param 			groupDetail 	body 	dto.UpdateGroupRequest 	true  "group detail"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateGroup(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.UpdateGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	group, err := h.repo.Get(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed fetch group")
	}

	if !group.IsOwner(userID) {
		return apperror.Forbidden("no permission", nil)
	}

	if err := group.Update(req.ProfilePicture, req.Name, userID); err != nil {
		return err
	}

	if err := h.repo.Update(group); err != nil {
		return errors.Wrap(err, "failed update group")
	}

	return c.SendStatus(fiber.StatusOK)
}
