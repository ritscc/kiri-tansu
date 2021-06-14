package model

import (
	"time"
)

type PasswordUser struct {
	UserID int `gorm:"primaryKey"`
	MailAddress string
	Password string
	IsLook bool
	CreatedAt time.Time
	UpdatedBy int
}
