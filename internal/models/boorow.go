package models

import (
	"gorm.io/gorm"
	"time"
)

type Borrow struct {
	gorm.Model
	UserID     uint
	BookID     uint
	BorrowDate time.Time
	ReturnDate *time.Time
}
