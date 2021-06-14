package model

type User struct {
	ID int `gorm:"primaryKey;autoIncrement:true"`
	Nickname string
	Role int
}
