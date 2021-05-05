package model

import (
	"time"
)

type Reservation struct {
	UserID int
	ItemID int
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
