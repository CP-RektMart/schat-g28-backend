package dto

import "github.com/CP-RektMart/schat-g28-backend/internal/model"

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type LoginRequest struct {
	IDToken string `json:"idToken" validate:"required"`
}

type LoginResponse struct {
	TokenResponse
	User UserResponse `json:"user"`
}

func ToTokenResponse(token model.Token) TokenResponse {
	return TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Exp:          token.Exp,
	}
}
