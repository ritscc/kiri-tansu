package model

import (
	"time"
)

type Lending struct {
	User_id int
	Item_id int
	Returned_at *time.Time
	Created_at *time.Time
	Created_by int
	Updated_at *time.Time
	Updated_by int
}
