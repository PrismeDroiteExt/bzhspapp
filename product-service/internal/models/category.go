package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	Name          string `gorm:"not null"`
	PictureUrl    string `gorm:"not null"`
	Description   string `gorm:"not null"`
	SubCategoryID *uint  `gorm:"default:null"`
}
