package entities

import "gorm.io/gorm"

type ReceiptItem struct {
	gorm.Model
	ProductCode string `gorm:"type:varchar(12)"`
	Name        string `gorm:"type:varchar(255); not null"`
	UnitPrice   uint32 `gorm:"not null"`
	Quantity    uint32 `gorm:"not null"`
	Origin      bool   `gorm:"not null"`
	ReceiptID   uint
}
