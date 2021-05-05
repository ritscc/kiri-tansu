package model

import (
	"time"
)

type PasswordResetUser struct {
	UserID int `gorm:"primaryKey"`
	ResetHash string
	ExpiredAt time.Time
}
