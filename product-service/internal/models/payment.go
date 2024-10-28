package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	OrderID       uint   `gorm:"not null"`
	Order         Order  `gorm:"foreignKey:OrderID"`
	UserID        uint   `gorm:"not null"`
	User          User   `gorm:"foreignKey:UserID"`
	Status        string `gorm:"not null"`
	PaymentMethod string `gorm:"not null"`
	TransactionId string `gorm:"not null"`
}
