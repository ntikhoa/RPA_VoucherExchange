package receipts

import (
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	TransactionID   string `gorm:"type:varchar(10)"`
	TransactionDate string `gorm:"type:varchar(20)"`
	StatusID        uint
	Status          ReceiptStatus `gorm:"foreignKey:StatusID"`
}
