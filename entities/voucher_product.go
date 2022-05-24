package entities

import (
	"time"

	"gorm.io/gorm"
)

type VoucherProduct struct {
	VoucherID   uint `gorm:"primaryKey"`
	ProductID   uint `gorm:"primaryKey"`
	QuantityMin uint `gorm:"not null"`
	QuantityMax uint `gorm:"not null"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
