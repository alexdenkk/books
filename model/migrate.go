package model

import (
	"gorm.io/gorm"
)

// Migrate - function for automigrating models to db
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Genre{})
}
