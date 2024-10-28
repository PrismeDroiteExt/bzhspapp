package tests

import (
	"errors"
	"reflect"
	"testing"

	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/models"
	"bzhspback.fr/breizhsport/internal/services"
	"gorm.io/gorm"
)

type MockBrandRepository struct {
	GetAllBrandsFunc func() ([]models.Brand, error)
	GetBrandByIDFunc func(id uint) (models.Brand, error)
}

func (m *MockBrandRepository) GetAllBrands() ([]models.Brand, error) {
	return m.GetAllBrandsFunc()
}

func (m *MockBrandRepository) GetBrandByID(id uint) (models.Brand, error) {
	return m.GetBrandByIDFunc(id)
}

func TestGetAllBrands(t *testing.T) {
	mockRepo := &MockBrandRepository{
		GetAllBrandsFunc: func() ([]models.Brand, error) {
			return []models.Brand{
				{Model: gorm.Model{ID: 1}, Name: "Brand A"},
				{Model: gorm.Model{ID: 2}, Name: "Brand B"},
			}, nil
		},
	}

	service := services.NewBrandService(mockRepo)

	expected := []dto.BrandResponse{
		{ID: 1, Name: "Brand A"},
		{ID: 2, Name: "Brand B"},
	}

	result, err := service.GetAllBrands()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetBrandByID(t *testing.T) {
	mockRepo := &MockBrandRepository{
		GetBrandByIDFunc: func(id uint) (models.Brand, error) {
			if id == 1 {
				return models.Brand{Model: gorm.Model{ID: 1}, Name: "Brand A"}, nil
			}
			return models.Brand{}, errors.New("brand not found")
		},
	}

	service := services.NewBrandService(mockRepo)

	// Test existing brand
	expected := dto.BrandResponse{ID: 1, Name: "Brand A"}

	result, err := service.GetBrandByID(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test non-existing brand
	_, err = service.GetBrandByID(2)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
