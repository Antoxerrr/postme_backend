package model

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	// Author
	Title   string `gorm:"type:varchar(128);not null"`
	Content string `gorm:"not null"`
	// Tags мб
}
