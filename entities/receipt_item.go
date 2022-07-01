package entities

import "gorm.io/gorm"

type ReceiptItem struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255); not null"`
	Amount    uint   `gorm:"not null"`
	ReceiptID uint
}
