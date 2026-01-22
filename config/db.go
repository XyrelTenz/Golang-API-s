package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/model"
)

var (
	DB *gorm.DB
)

func Database() {

	var err error

	DB, err := gorm.Open(sqlite.Open("backend.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("database not connected")

	}

	DB.AutoMigrate(&model.User{})
}
