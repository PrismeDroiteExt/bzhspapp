package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model

	CartID    uint    `gorm:"not null"`
	Cart      Cart    `gorm:"foreignKey:CartID"`
	ProductID uint    `gorm:"not null"`
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  uint    `gorm:"not null"`
}
