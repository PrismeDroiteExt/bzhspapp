package main

import (
	"log"

	"bzhspback.fr/breizhsport/internal/api/v1"
	"bzhspback.fr/breizhsport/internal/database"
	"github.com/gin-gonic/gin"
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
	r.Run(":8081")
}
