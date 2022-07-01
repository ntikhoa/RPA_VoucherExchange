package entities

import (
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	TransactionID string `gorm:"type:varchar(10); not null; index; UNIQUE"`
	StatusID      uint
	Status        CensorStatus   `gorm:"foreignKey:StatusID"`
	ReceiptItem   []ReceiptItem  `gorm:"foreignKey:ReceiptID"`
	Voucher       []Voucher      `gorm:"many2many:receipt_voucher"`
	ReceiptImage  []ReceiptImage `gorm:"foreignKey:ReceiptID"`
	Customer      Customer       `gorm:"foreignKey:ReceiptID"`
	AccountID     uint
	Account       Account `gorm:"foreignKey:AccountID"`
}
