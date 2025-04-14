package auth

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/dto"
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Login
// @Tags			auth
// @Router			/api/v1/auth/login [POST]
// @Param 			RequestBody 	body 	dto.LoginRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse[dto.LoginResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleLogin(c *fiber.Ctx) error {
	ctx := c.Context()

	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	OAuthUser, err := h.googleOauth.ValidateIDToken(ctx, req.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to validate id token")
	}

	user, err := h.repo.GetUserByEmail(OAuthUser.Email)
	if err != nil {
		user, err = model.NewUser(
			OAuthUser.Name,
			OAuthUser.Email,
			OAuthUser.ProfilePictureURL,
		)
		if err != nil {
			return err
		}

		user, err = h.repo.CreateUser(user)
		if err != nil {
			return err
		}
	}

	token, err := h.jwtService.GenerateAndStoreTokenPair(ctx, &user)
	if err != nil {
		return errors.Wrap(err, "failed to generate token pair")
	}

	result := dto.LoginResponse{
		TokenResponse: dto.ToTokenResponse(*token),
		User:          dto.ToUserResponse(user),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.LoginResponse]{
		Result: result,
	})
}
