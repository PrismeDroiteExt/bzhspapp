package controllers

import (
	"net/http"
	"strconv"

	"bzhspback.fr/breizhsport/internal/services"
	"github.com/gin-gonic/gin"
)

type BrandController struct {
	service *services.BrandService
}

// Constructor
func NewBrandController(service *services.BrandService) *BrandController {
	return &BrandController{service: service}
}

// GetAllBrands handles the GET request to retrieve all brands.
// It returns a JSON response with an array of brands or an error message.
//
// Route: GET /brands
// Response codes:
//   - 200 OK: Successfully retrieved brands
//   - 500 Internal Server Error: Failed to fetch brands
func (c *BrandController) GetAllBrands(ctx *gin.Context) {
	brands, err := c.service.GetAllBrands()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brands"})
		return
	}

	ctx.JSON(http.StatusOK, brands)
}

// GetBrandByID handles the GET request to retrieve a brand by its ID.
// It returns a JSON response with the brand details or an error message.
func (c *BrandController) GetBrandByID(ctx *gin.Context) {
	// get id parameter
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	brand, err := c.service.GetBrandByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch brand"})
		return
	}

	ctx.JSON(http.StatusOK, brand)
}
