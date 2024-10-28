package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model

	UserID      uint `gorm:"references:ID"`
	User        User `gorm:"foreignKey:UserID"`
	IsProcessed bool `gorm:"not null"`
}
