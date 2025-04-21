package model

import (
	"gorm.io/gorm"
)

type DirectMessage struct {
	gorm.Model
	Content string `gorm:"not null"`

	SenderID   uint `gorm:"not null"`
	ReceiverID uint `gorm:"not null"`
	Sender     User `gorm:"foreignKey:SenderID"`
	Receiver   User `gorm:"foreignKey:ReceiverID"`
}

type GroupMessage struct {
	gorm.Model
	Content string `gorm:"not null"`

	SenderID uint `gorm:"not null"`
	GroupID  uint `gorm:"not null"`
	Sender   User `gorm:"foreignKey:SenderID"`
	Group    Group `gorm:"foreignKey:GroupID"`
}
