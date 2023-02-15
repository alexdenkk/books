package model

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Review{})
	db.AutoMigrate(&Genre{})
}
