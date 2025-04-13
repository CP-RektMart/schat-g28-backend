package oauth

import (
	"context"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
)

type OAuth interface {
	ValidateIDToken(c context.Context, idToken string) (model.User, error)
}
