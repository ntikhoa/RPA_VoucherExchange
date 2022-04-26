package entities

import "gorm.io/gorm"

type GiftItem struct {
	gorm.Model
	GiftName  string `gorm:"type:varchar(255); not null"`
	VoucherID uint
}
