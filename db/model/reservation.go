package model

import (
	"time"
)

type Reservation struct {
	User_id int
	Item_id int
	Created_at *time.Time
	Created_by int
	Updated_at *time.Time
	Updated_by int
}
