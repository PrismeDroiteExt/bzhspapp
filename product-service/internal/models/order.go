package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	CartID uint   `gorm:"not null"`
	Cart   Cart   `gorm:"foreignKey:CartID"`
	Status string `gorm:"not null"`
}
