package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			get groups
// @Tags			groups
// @Router			/api/v1/groups [GET]
// @Security		ApiKeyAuth
// @Success			200 {object} 	dto.HttpListResponse[dto.GroupListResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetGroups(c *fiber.Ctx) error {
	group, err := h.repo.Get()
	if err != nil {
		return errors.Wrap(err, "failed fetch user")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.GroupListResponse]{
		Result: dto.ToGroupListsResponse(group),
	})
}
