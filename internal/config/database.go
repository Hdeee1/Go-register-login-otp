package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Hdeee1/go-register-login-otp/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
	fmt.Println("🚀 Database connected successfully!")
}

func InitDatabase() {
	if DB == nil {
		ConnectDatabase()
	}

	err := DB.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database!")
	}

	fmt.Println("🚀 Migration successful!")
}