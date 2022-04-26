package entities

import "gorm.io/gorm"

type VoucherItem struct {
	gorm.Model
	ItemName    string `gorm:"type:varchar(255); not null"`
	QuantityMin int    `gorm:"not null"`
	QuantityMax int    `gorm:"not null"`
	VoucherID   uint
}
