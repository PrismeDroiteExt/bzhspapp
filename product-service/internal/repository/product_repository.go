package repository

import (
	"fmt"
	"strings"

	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/models"
	"gorm.io/gorm"
)

// ProductRepository struct represents a repository for managing product data.
type ProductRepository struct {
	DB *gorm.DB
}

// Constructor for the ProductRepository struct.
func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// Get Product by ID
func (r *ProductRepository) GetProductByID(id uint) (models.Product, error) {
	var product models.Product

	result := r.DB.Preload("Category").Preload("Brand").First(&product, id)

	return product, result.Error
}

// Get all products for a category
func (r *ProductRepository) GetProductsByCategoryID(categoryID uint, filters []dto.Filter) ([]models.Product, error) {
	var products []models.Product

	query := `
		WITH RECURSIVE category_tree AS (
			SELECT id FROM categories WHERE id = ?
			UNION ALL
			SELECT c.id FROM categories c
			INNER JOIN category_tree ct ON c.sub_category_id = ct.id
		)
		SELECT DISTINCT p.* FROM products p
		INNER JOIN category_tree ct ON p.category_id = ct.id
		WHERE 1=1
	`
	args := []interface{}{categoryID}

	for _, filter := range filters {
		switch filter.Field {
		case "brand_id":
			query += " AND p.brand_id IN (?)"
			args = append(args, filter.Value)
		case "min_price":
			query += " AND p.price >= ?"
			args = append(args, filter.Value)
		case "max_price":
			query += " AND p.price <= ?"
			args = append(args, filter.Value)
		case "colors", "sizes":
			query += fmt.Sprintf(" AND string_to_array(p.%s, ',') && string_to_array(?, ',')", filter.Field)
			if str, ok := filter.Value.(string); ok {
				args = append(args, str)
			} else {
				return nil, fmt.Errorf("invalid type for %s filter: expected string", filter.Field)
			}
		}
	}

	// TODO: implement the pagination et the limit

	result := r.DB.Raw(query, args...).Preload("Brand").Preload("Category").Find(&products)

	return products, result.Error
}

// Get all recommended products
func (r *ProductRepository) GetRecommendedProducts() ([]models.Product, error) {
	var products []models.Product
	// recommended = discount > 0
	result := r.DB.Preload("Brand").Preload("Category").Where("discount > 0").Find(&products)

	return products, result.Error
}

// Get all products for a title or a description, with filters on the products
func (r *ProductRepository) GetProductsByNameOrDescription(titleOrDescription string, filters []dto.Filter) ([]models.Product, error) {
	var products []models.Product

	// Split the search query into words
	words := strings.Fields(titleOrDescription)
	tsquery := make([]string, len(words))
	for i, word := range words {
		tsquery[i] = fmt.Sprintf("%s:*", word)
	}

	query := `
		SELECT * FROM products WHERE 
		to_tsvector('english', title || ' ' || description) @@ to_tsquery('english', ?)
	`
	args := []interface{}{strings.Join(tsquery, " & ")}

	for _, filter := range filters {
		switch filter.Field {
		case "category_id":
			query += " AND category_id IN (?)"
			args = append(args, filter.Value)
		case "brand_id":
			query += " AND brand_id IN (?)"
			args = append(args, filter.Value)
		case "min_price":
			query += " AND price >= ?"
			args = append(args, filter.Value)
		case "max_price":
			query += " AND price <= ?"
			args = append(args, filter.Value)
		case "colors", "sizes":
			query += fmt.Sprintf(" AND string_to_array(%s, ',') && string_to_array(?, ',')", filter.Field)
			if str, ok := filter.Value.(string); ok {
				args = append(args, str)
			} else {
				return nil, fmt.Errorf("invalid type for %s filter: expected string", filter.Field)
			}
		}
	}

	result := r.DB.Raw(query, args...).Preload("Brand").Preload("Category").Find(&products)

	return products, result.Error
}
