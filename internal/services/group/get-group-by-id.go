package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			get group
// @Tags			groups
// @Router			/api/v1/groups/{groupID} [GET]
// @Security		ApiKeyAuth
// @Param 			groupID 	path 	uint 	true  "group id"
// @Success			200 {object} 	dto.HttpResponse[dto.GroupDetailResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetByID(c *fiber.Ctx) error {
	// ctx := c.UserContext()
	// userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	// if err != nil {
	// 	return errors.Wrap(err, "failed get userID from context")
	// }

	var req dto.GetGroupReqest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	group, err := h.repo.GetByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed fetch user")
	}

	// if !group.IsOwner(userID) && !group.IsMember(userID) {
	// 	return apperror.Forbidden("user not in the group", nil)
	// }

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.GroupDetailResponse]{
		Result: dto.ToGroupDetailReponse(group),
	})
}
