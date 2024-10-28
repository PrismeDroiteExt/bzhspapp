package services

import (
	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/repository"
	"github.com/jinzhu/copier"
)

// BrandService struct represents a service for managing brand data.
type BrandService struct {
	repo repository.BrandRepository
}

// Constructor
func NewBrandService(repo repository.BrandRepository) *BrandService {
	return &BrandService{repo: repo}
}

// GetAllBrands retrieves all brands from the database and returns them as a slice of BrandResponse.
func (s *BrandService) GetAllBrands() ([]dto.BrandResponse, error) {
	brands, err := s.repo.GetAllBrands()

	if err != nil {
		return nil, err
	}

	var brandDTOs []dto.BrandResponse
	err = copier.Copy(&brandDTOs, &brands)

	if err != nil {
		return nil, err
	}

	return brandDTOs, nil
}

// Get Brand by ID
func (s *BrandService) GetBrandByID(id uint) (dto.BrandResponse, error) {
	brand, err := s.repo.GetBrandByID(id)

	if err != nil {
		return dto.BrandResponse{}, err
	}

	return dto.BrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
	}, nil
}
