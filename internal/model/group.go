package model

import (
	"errors"

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
		return Group{}, err
	}

	return g, nil
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
	return g.Owner.ID == id
}

func (g *Group) IsMember(id uint) bool {
	_, found := lo.Find(g.Members, func(m User) bool {
		return m.ID == id
	})
	return found
}
