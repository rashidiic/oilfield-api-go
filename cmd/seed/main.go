package main

import (
	"fmt"
	"log"
	"time"

	"example.com/oilfield-api-go-two/internal/db"
	"example.com/oilfield-api-go-two/internal/models"

	"gorm.io/gorm"
)

func main() {
	database, err := db.InitDB("data/app.db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(database); err != nil {
		log.Fatal(err)
	}

	if err := Seed(database); err != nil {
		log.Fatal(err)
	}

	var fields, wells, sensors, readings int64
	database.Model(&models.OilField{}).Count(&fields)
	database.Model(&models.Well{}).Count(&wells)
	database.Model(&models.Sensor{}).Count(&sensors)
	database.Model(&models.ProductionReading{}).Count(&readings)

	fmt.Printf("OilFields=%d, Wells=%d, Sensors=%d, Readings=%d\n", fields, wells, sensors, readings)

	fmt.Println("✅ Week 2 seed completed")
}

func Seed(db *gorm.DB) error {
	db.Exec("DELETE FROM production_readings")
	db.Exec("DELETE FROM sensors")
	db.Exec("DELETE FROM wells")
	db.Exec("DELETE FROM oil_fields")

	parse := func(s string) time.Time {
		t, _ := time.Parse("2006-01-02", s)
		return t
	}

	fields := []models.OilField{
		{Name: "Caspian Ridge", Location: "Caspian Sea", OperatorCompany: "BlueWave Oil", StartDate: parse("2020-09-05")},
		{Name: "Absheron Onshore", Location: "Absheron", OperatorCompany: "GreenRock", StartDate: parse("2017-03-12")},
	}

	if err := db.Create(&fields).Error; err != nil {
		return err
	}

	for _, f := range fields {
		well := models.Well{
			OilFieldID: f.ID,
			Name:       f.Name + " Well A",
			Status:     models.WellActive,
			DrillDate:  parse("2019-01-01"),
			DepthM:     3200,
		}
		db.Create(&well)

		sensor := models.Sensor{
			WellID:      well.ID,
			SensorType:  models.SensorPressure,
			InstallDate: parse("2022-01-01"),
			IsActive:    true,
		}
		db.Create(&sensor)

		for i := 0; i < 5; i++ {
			db.Create(&models.ProductionReading{
				SensorID:  sensor.ID,
				Timestamp: time.Date(2025, 1, 1, i, 0, 0, 0, time.UTC),
				Value:     120 + float64(i)*1.5,
				Unit:      "bar",
			})
		}
	}

	return nil
}
