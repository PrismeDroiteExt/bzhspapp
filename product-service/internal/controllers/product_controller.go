package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"bzhspback.fr/breizhsport/internal/dto"
	"bzhspback.fr/breizhsport/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *services.ProductService
}

// Constructor
func NewProductController(service *services.ProductService) *ProductController {
	return &ProductController{service: service}
}

// GetProductByID godoc
// @Summary      Get product by ID
// @Description  Retrieves a product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  dto.ProductResponse
// @Failure      400  {object}  object{error=string}
// @Failure      500  {object}  object{error=string}
// @Router       /products/{id} [get]
func (c *ProductController) GetProductByID(ctx *gin.Context) {
	// get id parameter
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	product, err := c.service.GetProductByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// GetProductsByCategoryID godoc
// @Summary      Get products by category ID
// @Description  Retrieves all products for a specific category with optional filters
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        category_id  path      int     true   "Category ID"
// @Param        brand       query     string  false  "Brand IDs (comma-separated)"
// @Param        min_price   query     number  false  "Minimum price"
// @Param        max_price   query     number  false  "Maximum price"
// @Param        colors      query     string  false  "Colors (comma-separated)"
// @Param        sizes       query     string  false  "Sizes (comma-separated)"
// @Success      200  {array}   dto.ProductResponse
// @Failure      400  {object}  object{error=string}
// @Failure      500  {object}  object{error=string}
// @Router       /products/category/{category_id} [get]
func (c *ProductController) GetProductsByCategoryID(ctx *gin.Context) {
	categoryID, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	// Parse query parameters
	var filterRequest dto.FilterRequest
	filterRequest.Filters = make([]dto.Filter, 0)

	// Handle brand
	if brandIDsStr := ctx.Query("brand"); brandIDsStr != "" {
		brandIDs := strings.Split(brandIDsStr, ",")
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "brand_id", Value: brandIDs})
	}

	// Handle min_price
	if minPrice := ctx.Query("min_price"); minPrice != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "min_price", Value: minPrice})
	}

	// Handle max_price
	if maxPrice := ctx.Query("max_price"); maxPrice != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "max_price", Value: maxPrice})
	}

	// Handle colors
	if colors := ctx.Query("colors"); colors != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "colors", Value: colors})
	}

	// Handle sizes
	if sizes := ctx.Query("sizes"); sizes != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "sizes", Value: sizes})
	}

	products, err := c.service.GetProductsByCategoryID(uint(categoryID), filterRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetRecommendedProducts godoc
// @Summary      Get recommended products
// @Description  Retrieves a list of recommended products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   dto.ProductResponse
// @Failure      500  {object}  object{error=string}
// @Router       /products/recommended [get]
func (c *ProductController) GetRecommendedProducts(ctx *gin.Context) {
	products, err := c.service.GetRecommendedProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch recommended products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}

// GetProductsByNameOrDescription godoc
// @Summary      Search products
// @Description  Search products by name or description with optional filters
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        query       query     string  true   "Search query"
// @Param        category    query     int     false  "Category ID"
// @Param        brand       query     string  false  "Brand IDs (comma-separated)"
// @Param        min_price   query     number  false  "Minimum price"
// @Param        max_price   query     number  false  "Maximum price"
// @Param        colors      query     string  false  "Colors (comma-separated)"
// @Param        sizes       query     string  false  "Sizes (comma-separated)"
// @Success      200  {array}   dto.ProductResponse
// @Failure      400  {object}  object{error=string}
// @Failure      500  {object}  object{error=string}
// @Router       /products/search [get]
func (c *ProductController) GetProductsByNameOrDescription(ctx *gin.Context) {
	nameOrDescription := ctx.Query("query")
	if nameOrDescription == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// Parse other filters as before
	var filterRequest dto.FilterRequest
	filterRequest.Filters = make([]dto.Filter, 0)

	// Handle category
	if categoryID, err := strconv.Atoi(ctx.Query("category")); err == nil {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "category_id", Value: categoryID})
	}

	// Handle brand
	if brandIDsStr := ctx.Query("brand"); brandIDsStr != "" {
		brandIDs := strings.Split(brandIDsStr, ",")
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "brand_id", Value: brandIDs})
	}

	// Handle min_price
	if minPrice := ctx.Query("min_price"); minPrice != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "min_price", Value: minPrice})
	}

	// Handle max_price
	if maxPrice := ctx.Query("max_price"); maxPrice != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "max_price", Value: maxPrice})
	}

	// Handle colors
	if colors := ctx.Query("colors"); colors != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "colors", Value: colors})
	}

	// Handle sizes
	if sizes := ctx.Query("sizes"); sizes != "" {
		filterRequest.Filters = append(filterRequest.Filters, dto.Filter{Field: "sizes", Value: sizes})
	}

	products, err := c.service.GetProductsByNameOrDescription(nameOrDescription, filterRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	ctx.JSON(http.StatusOK, products)
}
