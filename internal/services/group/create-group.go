package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			create group
// @Tags			groups
// @Router			/api/v1/groups [POST]
// @Security		ApiKeyAuth
// @Param 			groupDetail 	body 	dto.CreateGroupRequest 	true  "group detail"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreateGroup(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed get userID from context")
	}

	var req dto.CreateGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	group, err := model.NewGroup(req.ProfilePictureURL, req.Name, userID, req.MemberIDs)
	if err != nil {
		return apperror.BadRequest(err.Error(), err)
	}

	if err := h.repo.Create(group); err != nil {
		return apperror.BadRequest("user not exist", err)
	}

	return c.SendStatus(fiber.StatusOK)
}
