package model

import (
	"time"
)

type Lending struct {
	UserID int `gorm:"primaryKey"`
	ItemID int `gorm:"primaryKey"`
	ReturnedAt time.Time
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
