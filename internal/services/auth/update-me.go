package auth

import (
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"

	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Update me
// @Description		Update user's profile
// @Tags			user
// @Router			/api/v1/me [PATCH]
// @Security		ApiKeyAuth
// @Param 			RequestBody 	body 	dto.UserUpdateRequest 	true 	"request request"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateMe(c *fiber.Ctx) error {
	ctx := c.UserContext()
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UserUpdateRequest)
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	u, err := h.repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	err = u.Update(req.Name, req.Email, req.ProfilePictureURL, req.Color)
	if err != nil {
		return err
	}

	err = h.repo.UpdateUser(u)
	if err != nil {
		return err
	}

	response := dto.ToUserResponse(u)

	return c.JSON(dto.HttpResponse[dto.UserResponse]{
		Result: response,
	})
}
