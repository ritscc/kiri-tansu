package model

import (
	"time"
)

type Password_user struct {
	User_id int
	Mail_address string
	Password string
	Is_look bool
	Created_at *time.Time
	Updated_by int
}
