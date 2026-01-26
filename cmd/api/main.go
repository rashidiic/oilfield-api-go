package main

import (
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example.com/oilfield-api-go-two/docs"

	"example.com/oilfield-api-go-two/internal/db"
	"example.com/oilfield-api-go-two/internal/mock"
)

// @title           Oilfield API
// @version         1.0
// @description     Week 1: Mock CRUD | Week 2: Oilfield schema + seed
// @BasePath        /
func main() {
	// --- Init DB ---
	database, err := db.InitDB("data/app.db")
	if err != nil {
		log.Fatal(err)
	}

	// --- AutoMigrate (Week 1 + Week 2) ---
	if err := db.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	// --- Gin ---
	r := gin.Default()

	// --- API group ---
	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	// --- Mock CRUD (Week 1) ---
	mock.RegisterRoutes(api, database)

	// --- Swagger ---
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// --- Run server ---
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
