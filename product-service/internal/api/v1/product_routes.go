package api

import (
	"bzhspback.fr/breizhsport/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Parameters:
//   - router: A pointer to the gin.RouterGroup where the product routes will be added.
//   - controller: A pointer to the ProductController that handles brand-related operations.
func SetupProductRoutes(router *gin.RouterGroup, controller *controllers.ProductController) {
	productRoutes := router.Group("/products") // Group all product routes under /brands
	{
		// GET /brands - Retrieves all brands
		productRoutes.GET("/:id", controller.GetProductByID)
		productRoutes.GET("/category/:category_id", controller.GetProductsByCategoryID)
		productRoutes.GET("/recommended", controller.GetRecommendedProducts)
		productRoutes.GET("/search", controller.GetProductsByNameOrDescription)
	}
}
