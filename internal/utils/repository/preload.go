package repository

import "gorm.io/gorm"

func AccumulatePreload(db *gorm.DB, preload ...string) *gorm.DB {
	for _, p := range preload {
		db = db.Preload(p)
	}
	return db
}
