package receipts

import "gorm.io/gorm"

type ReceiptItem struct {
	gorm.Model
	ProductCode string `gorm:"type:varchar(12)"`
	Name        string `gorm:"type:varchar(255)"`
	UnitPrice   uint32
	Quantity    uint32
	ReceiptID   uint
	Receipt     Receipt `gorm:"foreignKey:ReceiptID"`
}
