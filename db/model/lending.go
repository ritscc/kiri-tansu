package model

import (
	"time"
)

type Lending struct {
	UserID int
	ItemID int
	ReturnedAt time.Time
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
