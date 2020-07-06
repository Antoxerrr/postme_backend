package model

import (
	"github.com/jinzhu/gorm"
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&Test{}, &Post{})
}
