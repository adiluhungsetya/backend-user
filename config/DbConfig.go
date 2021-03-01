package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"gorm.io/driver/postgres"

)

type DataSource struct {
	DbUser string
	DbPass string
	DbHost string
	DbName string
	DbPort string
	DbEngine string
}

func Connect() *gorm.DB {
	DB := DataSource{
		DbUser: viper.GetString("DB_USER"),
		DbPass: viper.GetString("DB_PASSWORD"),
		DbHost: viper.GetString("DB_HOST"),
		DbName: viper.GetString("DB_NAME"),
		DbPort: viper.GetString("DB_PORT"),
		DbEngine: viper.GetString("DB_ENGINE"),
	}

	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", DB.DbHost, DB.DbUser, DB.DbPass, DB.DbName, DB.DbPort)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
