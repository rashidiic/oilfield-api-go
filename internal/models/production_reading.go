package models

import "time"

type ProductionReading struct {
	ID        uint      `gorm:"primaryKey"`
	SensorID  uint      `gorm:"not null;index"`
	Timestamp time.Time `gorm:"not null;index"`
	Value     float64   `gorm:"not null"`
	Unit      string    `gorm:"not null"`

	Sensor Sensor `gorm:"foreignKey:SensorID;references:ID"`
}
