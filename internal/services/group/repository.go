package group

import (
	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/internal/utils/repository"
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

func (r *Repository) Get() ([]model.Group, error) {
	var g []model.Group

	if err := r.db.Find(&g).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Group{}, apperror.NotFound("groups not found", err)
		}
		return []model.Group{}, err
	}

	return g, nil
}

func (r *Repository) GetByID(id uint, preload ...string) (model.Group, error) {
	var g model.Group

	db := repository.AccumulatePreload(r.db, preload...)

	if err := db.Preload("Members").Preload("Messages").First(&g, id).Error; err != nil {
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

	return r.db.Model(&group).Omit("Members.*").Association("Members").Append(&user)
}

func (r *Repository) LeaveGroup(groupID, userID uint) error {
	group := model.Group{Model: gorm.Model{ID: groupID}}
	user := model.User{Model: gorm.Model{ID: userID}}

	return r.db.Model(&group).Association("Members").Delete(&user)
}
