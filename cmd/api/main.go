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

// @title           Oilfield API (Week 1 Mock CRUD)
// @version         1.0
// @description     Week 1: Gin + GORM + SQLite + Swagger + Mock CRUD
// @BasePath        /
func main() {
	database, err := db.InitDB("data/app.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})
	}

	mock.RegisterRoutes(api, database)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
