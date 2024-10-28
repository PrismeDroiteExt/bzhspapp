package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	FirstName    string
	LastName     string
	Role         string
	LastLogin    time.Time
	RefreshToken string
	Active       bool
}
