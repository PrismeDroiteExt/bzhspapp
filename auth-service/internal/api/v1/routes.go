package api

import (
	"github.com/gin-gonic/gin"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/config"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/controllers"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/database"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/middleware"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/repository"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/services"
)

func InitRoutes(r *gin.Engine) {
	// Initialize dependencies
	db := database.GetDB()
	cfg := config.LoadConfig()
	repo := repository.NewAuthRepository(db)
	service := services.NewAuthService(repo, cfg)
	controller := controllers.NewAuthController(service)

	auth := r.Group("/api/v1/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.POST("/refresh", controller.RefreshToken)

		// Protected routes
		protected := auth.Group("/")
		protected.Use(middleware.JWTAuthMiddleware(cfg))
		{
			protected.POST("/logout", controller.Logout)
			protected.GET("/me", controller.GetProfile)
			protected.PUT("/me", controller.UpdateProfile)
		}
	}
}
