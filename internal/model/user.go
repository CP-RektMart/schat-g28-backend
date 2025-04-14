package model

import (
	"net/mail"

	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string `gorm:"not null"`
	Email             string `gorm:"not null;unique"`
	ProfilePictureURL string
}

func NewUser(name, email, profilePictureURL string) (User, error) {
	u := User{
		Name:              name,
		Email:             email,
		ProfilePictureURL: profilePictureURL,
	}

	if err := u.Valid(); err != nil {
		return User{}, nil
	}

	return u, nil
}

func (u *User) Update(name, email, profilePictureURL *string) error {
	if name != nil {
		u.Name = *name
	}

	if email != nil {
		u.Email = *email
	}

	if profilePictureURL != nil {
		u.ProfilePictureURL = *profilePictureURL
	}

	return u.Valid()
}

func (u *User) Valid() error {
	if u.Name == "" {
		return apperror.BadRequest("name is required", nil)
	}

	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return apperror.BadRequest("email is invalid", nil)
	}

	return nil
}

