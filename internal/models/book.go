package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null"`
	Quantity int    `gorm:"not null"`
	Code     string `gorm:"type:varchar(50);unique;not null"`
}
