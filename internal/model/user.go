package model

type User struct {
	ID       int64  `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
}
