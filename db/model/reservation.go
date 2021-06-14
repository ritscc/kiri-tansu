package model

import (
	"time"
)

type Reservation struct {
	UserID int `gorm:"primaryKey"`
	ItemID int `gorm:"primaryKey"`
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
