package model

import (
	"github.com/google/uuid"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	Exp          int64
}

type CachedTokens struct {
	AccessUID  uuid.UUID
	RefreshUID uuid.UUID
}
