package entities

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Name           string `gorm:"type:varchar(255)"`
	Description    string `gorm:"type:varchar(255)"`
	TotalPriceMin  uint   `gorm:"not null"`
	TotalPriceMax  uint   `gorm:"not null"`
	Total          uint   `gorm:"not null"`
	Remaining      uint   `gorm:"not null"`
	Published      bool   `gorm:"not null"`
	VoucherProduct []VoucherProduct
	Gift           []Gift `gorm:"foreignKey:VoucherID"`
	ProviderID     uint   `json:"-"`
}
