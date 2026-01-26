package models

import "time"

type WellStatus string

const (
	WellActive    WellStatus = "active"
	WellShutIn    WellStatus = "shut-in"
	WellAbandoned WellStatus = "abandoned"
)

func (s WellStatus) IsValid() bool {
	return s == WellActive || s == WellShutIn || s == WellAbandoned
}

type Well struct {
	ID         uint       `gorm:"primaryKey"`
	OilFieldID uint       `gorm:"not null;index"`
	Name       string     `gorm:"not null"`
	Status     WellStatus `gorm:"type:text;not null"`
	DrillDate  time.Time  `gorm:"not null"`
	DepthM     float64    `gorm:"not null"`

	OilField OilField `gorm:"foreignKey:OilFieldID;references:ID"`
	Sensors  []Sensor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
