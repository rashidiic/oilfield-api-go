package models

import "time"

type SensorType string

const (
	SensorPressure    SensorType = "pressure"
	SensorTemperature SensorType = "temperature"
	SensorFlowRate    SensorType = "flowrate"
)

func (t SensorType) IsValid() bool {
	return t == SensorPressure || t == SensorTemperature || t == SensorFlowRate
}

type Sensor struct {
	ID          uint       `gorm:"primaryKey"`
	WellID      uint       `gorm:"not null;index"`
	SensorType  SensorType `gorm:"type:text;not null"`
	InstallDate time.Time  `gorm:"not null"`
	IsActive    bool       `gorm:"not null;default:true"`

	Well     Well                `gorm:"foreignKey:WellID;references:ID"`
	Readings []ProductionReading `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
