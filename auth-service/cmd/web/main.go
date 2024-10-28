package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/api/v1"
	"github.com/prismedroiteext/breizhsport/auth-service/internal/database"
)

func main() {
	r := gin.Default()

	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize routes
	api.InitRoutes(r)

	// Start server
	r.Run(":8082")
}
