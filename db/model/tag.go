package model

import (
	"time"
)

type Tag struct {
	Id int
	Name string
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
