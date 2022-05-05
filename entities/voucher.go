package entities

import (
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Description   string     `gorm:"type:varchar(255)"`
	TotalPriceMin int        `gorm:"not null"`
	TotalPriceMax int        `gorm:"not null"`
	Published     bool       `gorm:"not null"`
	Product       []*Product `gorm:"many2many:voucher_products"`
	Gift          []*Gift    `gorm:"foreignKey:VoucherID"`
}
