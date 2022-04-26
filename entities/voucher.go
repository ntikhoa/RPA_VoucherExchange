package entities

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Description   string        `gorm:"type:varchar(255)"`
	TotalPriceMin int           `gorm:"not null"`
	TotalPriceMax int           `gorm:"not null"`
	Published     bool          `gorm:"not null"`
	VoucherItem   []VoucherItem `gorm:"foreignKey:VoucherID"`
	GiftItem      []GiftItem    `gorm:"foreignKey:VoucherID"`
	Receipt       []*Receipt    `gorm:"many2many:receipt_voucher"`
}
