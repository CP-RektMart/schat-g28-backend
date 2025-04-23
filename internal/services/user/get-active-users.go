package user

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// @Summary			get active users
// @Tags			user
// @Router			/api/v1/users/active [GET]
// @Security		ApiKeyAuth
// @Success			200 {object} 	dto.HttpListResponse[dto.UserResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetActiveUsers(c *fiber.Ctx) error {
	users, err := h.repo.Get()
	if err != nil {
		return errors.Wrap(err, "failed fetch user")
	}

	clients := h.chatService.GetActiveUsers()
	filtered := lo.Filter(users, func(u model.User, _ int) bool {
		_, ok := clients[u.ID]
		return ok
	})

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.UserResponse]{
		Result: dto.ToUsersReponse(filtered),
	})
}
