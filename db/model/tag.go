package model

import (
	"time"
)

type Tag struct {
	Id int
	Name string
	Created_at *time.Time
	Created_by int
	Updated_at *time.Time
	Updated_by int
}
