package entities

import "gorm.io/gorm"

type ReceiptImage struct {
	gorm.Model
	Url       string `gorm:"type:varchar(255); not null"`
	ReceiptID uint
}
