package db

import (
	"os"
	"path/filepath"

	"example.com/oilfield-api-go-two/internal/mock"
	"example.com/oilfield-api-go-two/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(sqlitePath string) (*gorm.DB, error) {
	dir := filepath.Dir(sqlitePath)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}

	database, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return database, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&mock.MockItem{},
		&models.OilField{},
		&models.Well{},
		&models.Sensor{},
		&models.ProductionReading{},
	)
}
