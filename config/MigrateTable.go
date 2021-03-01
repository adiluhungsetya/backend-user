package config

import (
	"gorm.io/gorm"
	"log"
	"user-go-test/model"
)

func MigrateTable(db *gorm.DB)  {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal(err)
	}
}