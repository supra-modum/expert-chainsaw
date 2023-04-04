package models

import (
	"gorm.io/gorm"
)

type Fundraising struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey;column:id"`
	Title       string  `gorm:"size:255;column:title"`
	Description string  `gorm:"size:1024;column:description"`
	CampaignURL string  `gorm:"size:255;column:url"`
	GoalAmount  float64 `gorm:"column:goal"`
	Closed      bool    `gorm:"column:closed"`
}
