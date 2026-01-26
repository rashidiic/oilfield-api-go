package models

import "time"

type OilField struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null" json:"name"`
	Location        string    `gorm:"not null" json:"location"`
	OperatorCompany string    `gorm:"not null" json:"operatorCompany"`
	StartDate       time.Time `gorm:"not null" json:"startDate"`

	Wells []Well `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
