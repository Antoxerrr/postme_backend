package model

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	// Author
	Text string `gorm:"type:varchar(512)"`
	Post *Post
}
