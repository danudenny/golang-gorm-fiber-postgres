package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	Username  string `gorm:"unique"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Users []User
