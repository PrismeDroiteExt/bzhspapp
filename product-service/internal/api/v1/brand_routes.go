package api

import (
	"bzhspback.fr/breizhsport/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Parameters:
//   - router: A pointer to the gin.RouterGroup where the brand routes will be added.
//   - controller: A pointer to the BrandController that handles brand-related operations.
func SetupBrandRoutes(router *gin.RouterGroup, controller *controllers.BrandController) {
	brandRoutes := router.Group("/brands") // Group all brand routes under /brands
	{
		// GET /brands - Retrieves all brands
		brandRoutes.GET("", controller.GetAllBrands)
		brandRoutes.GET("/:id", controller.GetBrandByID)

		// TODO: Add more brand-related routes here as needed
		// For example:
		// brandRoutes.GET("/:id", controller.GetBrandByID)
		// brandRoutes.POST("", controller.CreateBrand)
		// brandRoutes.PUT("/:id", controller.UpdateBrand)
		// brandRoutes.DELETE("/:id", controller.DeleteBrand)
	}
}
