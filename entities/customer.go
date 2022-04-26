package entities

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255)"`
	Phone     string `gorm:"type:varchar(12)"`
	ReceiptID uint
	Receipt   Receipt `gorm:"foreignKey:ReceiptID"`
}
