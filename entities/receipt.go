package entities

import (
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	TransactionID   string `gorm:"type:varchar(10); not null; index"`
	TransactionDate string `gorm:"type:varchar(20); not null"`
	StatusID        uint
	Status          ReceiptStatus `gorm:"foreignKey:StatusID"`
	ReceiptItem     []ReceiptItem `gorm:"foreignKey:ReceiptID"`
	Voucher         []Voucher     `gorm:"many2many:receipt_voucher"`
}
