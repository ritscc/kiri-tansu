package model

import (
	"time"
)

type LdapUser struct {
	UserID int `gorm:"primaryKey"`
	LdapID int
	CreatedAt time.Time
	CreatedBy int
	UpdatedAt time.Time
	UpdatedBy int
}
