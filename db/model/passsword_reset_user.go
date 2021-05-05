package model

import (
	"time"
)

type Password_reset_user struct {
	User_id int
	Reset_hash string
	Expired_at *time.Time
}
