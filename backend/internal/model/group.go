package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	name string `gorm:"not null"`
}