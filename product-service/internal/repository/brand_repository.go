package repository

import (
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

type BrandRepository interface {
	GetAllBrands() ([]models.Brand, error)
	GetBrandByID(id uint) (models.Brand, error)
}

// BrandRepositoryImpl struct represents a repository for managing brand data.
type BrandRepositoryImpl struct {
	DB *gorm.DB
}

// Constructor for the BrandRepositoryImpl struct.
func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &BrandRepositoryImpl{DB: db}
}

// GetAllBrands retrieves all brands from the database.
func (r *BrandRepositoryImpl) GetAllBrands() ([]models.Brand, error) {
	var brands []models.Brand
	result := r.DB.Find(&brands)
	return brands, result.Error
}

// Get Brand by ID
func (r *BrandRepositoryImpl) GetBrandByID(id uint) (models.Brand, error) {
	var brand models.Brand
	result := r.DB.First(&brand, id)
	return brand, result.Error
}
