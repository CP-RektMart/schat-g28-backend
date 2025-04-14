package model

import (
	"errors"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	URL     string `gorm:"not null"`
	Path    *string
	OwnerID uint `gorm:"not null"`
	Owner   User `gorm:"foreignKey:OwnerID"`
}

func NewFile(URL string, path *string, ownerID uint) (File, error) {
	f := File{
		URL:     URL,
		OwnerID: ownerID,
		Path:    path,
	}
	if err := f.Valid(); err != nil {
		return File{}, err
	}
	return f, nil
}

func (f *File) Valid() error {
	if f.URL == "" {
		return errors.New("URL cannot be empty")
	}

	if f.OwnerID == 0 {
		return errors.New("ownerID cannot be empty")
	}

	return nil
}

func (f *File) IsOwner(userID uint) bool {
	return f.OwnerID == userID
}
