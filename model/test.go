package model

import "github.com/jinzhu/gorm"

type Test struct {
	gorm.Model
	Text string
}
