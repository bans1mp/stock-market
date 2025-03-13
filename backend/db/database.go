package db

import (
	"backend/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	// Get database connection string from .env
	dsn := os.Getenv("DATABASE_URL")

	// Connect to MySQL using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
		return nil, err
	}

	db.AutoMigrate(&models.User{}, &models.UserBalance{}, &models.Trade{}, &models.Stock{}, &models.UserStock{})
	DB = db
	return db, nil
}
