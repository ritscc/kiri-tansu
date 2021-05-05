package model

import (
	"time"
)

type Item struct {
	ID int `gorm:"primaryKey;autoIncrement:true"`
	Name string
	Overview string
	Icon byte
	Type int
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
