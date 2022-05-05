package entities

import (
	"time"

	"gorm.io/gorm"
)

type VoucherProduct struct {
	VoucherID   uint  `gorm:"primaryKey"`
	ProductID   uint  `gorm:"primaryKey"`
	QuantityMin uint8 `gorm:"not null"`
	QuantityMax uint8 `gorm:"not null"`
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
