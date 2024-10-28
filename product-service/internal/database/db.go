package database

import (
	"fmt"
	"log"
	"os"

	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate
	err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{}, &models.Order{}, &models.Brand{}, &models.Cart{}, &models.CartItem{}, &models.Payment{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
