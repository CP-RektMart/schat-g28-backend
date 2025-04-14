package auth

import (
	"errors"
	"strings"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetUserByID(id uint) (model.User, error) {
	var u model.User
	err := r.db.Preload("Groups").Preload("Friends").First(&u, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, apperror.NotFound("user don't exist", err)
		}
		return model.User{}, err
	}
	return u, nil
}

func (r *Repository) GetUserByEmail(email string) (model.User, error) {
	var u model.User
	err := r.db.Where("email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, apperror.NotFound("user with this email not exist", err)
		}
		return model.User{}, err
	}
	return u, nil
}

func (r *Repository) CreateUser(u model.User) (model.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return model.User{}, apperror.BadRequest("email already exist", err)
		}
		return model.User{}, err
	}
	return u, nil
}

func (r *Repository) UpdateUser(u model.User) error {
	err := r.db.Save(&u).Error
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return apperror.BadRequest("email already exist", err)
		}
		return err
	}

	return nil
}

func (r *Repository) AddFriend(userID, friendID uint) error {
	user := model.User{Model: gorm.Model{ID: userID}}
	friend := model.User{Model: gorm.Model{ID: friendID}}

	return r.db.Model(&user).Association("Friends").Append(&friend)
}

func (r *Repository) UnFriend(userID, friendID uint) error {
	user := model.User{Model: gorm.Model{ID: userID}}
	friend := model.User{Model: gorm.Model{ID: friendID}}

	return r.db.Model(&user).Association("Friends").Delete(&friend)
}
