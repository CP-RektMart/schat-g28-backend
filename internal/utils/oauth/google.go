package oauth

import (
	"context"
	"fmt"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"google.golang.org/api/idtoken"
)

type GoogleConfig struct {
	ClientID string `env:"CLIENT_ID"`
}

type Google struct {
	config GoogleConfig
}

func NewGoogle(config GoogleConfig) OAuth {
	return &Google{
		config: config,
	}
}

func (g *Google) ValidateIDToken(c context.Context, idToken string) (model.User, error) {
	payload, err := idtoken.Validate(c, idToken, g.config.ClientID)
	if err != nil {
		return model.User{}, err
	}

	info := map[string]string{
		"name":    "",
		"email":   "",
		"picture": "",
	}

	for field := range info {
		value, ok := payload.Claims[field].(string)
		if !ok {
			return model.User{}, fmt.Errorf("%s not found in idToken claim", field)
		}
		info[field] = value
	}

	return model.User{
		Name:              info["name"],
		Email:             info["email"],
		ProfilePictureURL: info["picture"],
	}, nil
}
