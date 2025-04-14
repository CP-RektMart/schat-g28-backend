package model

import (
	"errors"

	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ProfilePictureURL *string
	Name              string         `gorm:"not null"`
	OwnerID           uint           `gorm:"not null"`
	Owner             User           `gorm:"foreignKey:OwnerID"`
	Members           []User         `gorm:"many2many:group_member"`
	Messages          []GroupMessage `gorm:"foreignKey:GroupID"`
}

func NewGroup(profilePicture *string, name string, ownerID uint) (Group, error) {
	g := Group{
		Name:              name,
		ProfilePictureURL: profilePicture,
		OwnerID:           ownerID,
	}
	if err := g.Valid(); err != nil {
		return Group{}, apperror.BadRequest("invalid input", err)
	}

	return g, nil
}

func (g *Group) Update(profilePicture, name *string, changerID uint) error {
	if !g.IsOwner(changerID) {
		return apperror.Forbidden("user not an owner of the group", nil)
	}

	if profilePicture != nil {
		g.ProfilePictureURL = profilePicture
	}

	if name != nil {
		g.Name = *name
	}

	if err := g.Valid(); err != nil {
		return apperror.BadRequest("invalid input", err)
	}

	return nil
}

func (g *Group) Valid() error {
	if g.Name == "" {
		return errors.New("name cannot be empty")
	}

	if g.OwnerID == 0 {
		return errors.New("owner id cannot be empty")
	}

	return nil
}

func (g *Group) IsOwner(id uint) bool {
	return g.OwnerID == id
}

func (g *Group) IsMember(id uint) bool {
	_, found := lo.Find(g.Members, func(m User) bool {
		return m.ID == id
	})
	return found
}
