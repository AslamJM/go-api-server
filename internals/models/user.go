package models

import "time"

type User struct {
	ID         uint    `gorm:"primaryKey"`
	Username   string  `gorm:"size:255"`
	Email      *string `gorm:"size:255;uniqueIndex"`
	Password   string  `gorm:"size:255"`
	Created_at time.Time
	Updated_at time.Time
}
