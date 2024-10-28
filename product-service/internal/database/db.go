package database

import (
	"fmt"
	"os"

	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("PRODUCT_DB_HOST"),
		os.Getenv("PRODUCT_DB_USER"),
		os.Getenv("PRODUCT_DB_PASSWORD"),
		os.Getenv("PRODUCT_DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to product database: %v", err)
	}

	// Auto migrate models
	return db.AutoMigrate(&models.Product{}, &models.Brand{}, &models.Category{})
}

func GetDB() *gorm.DB {
	return db
}
