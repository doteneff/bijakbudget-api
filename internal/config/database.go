package config

import (
	"fmt"
	"log"
	"os"

	"github.com/doteneff/bijakbudget-api/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using OS env vars")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database

	// Migrate the schema
	err = DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connection established and models migrated!")
}