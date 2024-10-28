package main

import (
	"bzhspback.fr/breizhsport/internal/api/v1"
	"bzhspback.fr/breizhsport/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.InitDB()
	db := database.GetDB()
	api.InitRoutes(r, db)

	r.Run()
}
