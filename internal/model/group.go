package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	ProfilePicture string `gorm:"not null"`
	Name           string `gorm:"not null"`
}
