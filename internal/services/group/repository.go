package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(g model.Group) error {
	return r.db.Debug().Omit("Members.*").Create(&g).Error
}

func (r *Repository) Get(id uint) (model.Group, error) {
	var g model.Group
	if err := r.db.Preload("Owner").Preload("Messages").Preload("Members").First(&g, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Group{}, apperror.NotFound("group not found", err)
		}
		return model.Group{}, err
	}
	return g, nil
}

func (r *Repository) Update(g model.Group) error {
	return r.db.Save(&g).Error
}

func (r *Repository) Delete(id uint) error {
	if err := r.db.Delete(&model.Group{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("group not found", err)
		}
		return err
	}
	return nil
}

func (r *Repository) JoinGroup(groupID, userID uint) error {
	group := model.Group{Model: gorm.Model{ID: groupID}}
	user := model.User{Model: gorm.Model{ID: userID}}

	return r.db.Model(&group).Association("Members").Append(&user)
}

func (r *Repository) LeaveGroup(groupID, userID uint) error {
	group := model.Group{Model: gorm.Model{ID: groupID}}
	user := model.User{Model: gorm.Model{ID: userID}}

	return r.db.Model(&group).Association("Members").Delete(&user)
}
