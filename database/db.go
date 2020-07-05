package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"postme/config"
)

var db *gorm.DB
var err error

func InitDB(config *config.Config) *gorm.DB {
	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB.PostgresHost, config.DB.PostgresUser,
		config.DB.PostgresDB, config.DB.PostgresPassword)
	db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}
