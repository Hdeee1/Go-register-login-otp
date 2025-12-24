package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	GetEnv("DB_HOST", "localhost"),
	GetEnv("DB_PORT", "5432"),
	GetEnv("DB_USER", "postgres"),
	GetEnv("DB_PASSWORD", ""),
	GetEnv("DB_NAME", ""))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect Database:" + err.Error())
	}
	DB = db
}