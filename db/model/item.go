package model

import (
	"time"
)

type Item struct {
	Id int
	Name string
	Overview string
	Icon byte
	Type int
	Created_at *time.Time
	Created_by int
	Updated_at *time.Time
	Updated_by int
}
