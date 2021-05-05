package model

import (
	"time"
)

type PasswordResetUser struct {
	UserID int
	ResetHash string
	ExpiredAt time.Time
}
