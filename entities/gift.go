package entities

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	GiftName   string `gorm:"type:varchar(255); not null"`
	VoucherID  uint
	ProviderID uint `json:"-"`
}
