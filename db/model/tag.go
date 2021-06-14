package model

import (
	"time"
)

type Tag struct {
	ID int `gorm:"primaryKey;autoIncrement:true"`
	Name string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
