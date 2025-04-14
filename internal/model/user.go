package model

import (
	"errors"
	"net/mail"

	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string `gorm:"not null"`
	Email             string `gorm:"not null;unique"`
	ProfilePictureURL string
	Friends           []User  `gorm:"many2many:friends"`
	Groups            []Group `gorm:"many2many:group_member"`
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

func (u *User) CanbeFriend(userID uint) error {
	if u.ID == userID {
		return errors.New("user can't be a friend with themself")
	}

	if u.IsFriend(userID) {
		return errors.New("already be a friend")
	}

	return nil
}

func (u *User) CanUnFriend(userID uint) error {
	if u.ID == userID {
		return errors.New("user can't unfriend themself")
	}

	if !u.IsFriend(userID) {
		return errors.New("not a friend")
	}

	return nil
}

func (u *User) IsFriend(userID uint) bool {
	_, found := lo.Find(u.Friends, func(f User) bool {
		return f.ID == userID
	})
	return found
}
