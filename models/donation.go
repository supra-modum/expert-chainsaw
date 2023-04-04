package models

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	gorm.Model
	ID            uint `gorm:"primaryKey"`
	UserID        uint `gorm:"index"`
	FundraisingID uint `gorm:"index"`
	Amount        float64
	Sent          bool `gorm:"not null;default:false"`
	CreatedAt     time.Time
}
