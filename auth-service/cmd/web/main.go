package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/prismedroiteext/breizhsport/auth-service/docs" // This is important!
	"github.com/prismedroiteext/breizhsport/auth-service/internal/api/v1"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/database"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           BreizhSport Auth Service API
// @version         1.0
// @description     Authentication and authorization service for BreizhSport application
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /api/v1/auth
// @schemes   http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	r := gin.Default()

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize routes
	api.InitRoutes(r)

	// Start server
	r.Run(":8082")
}
