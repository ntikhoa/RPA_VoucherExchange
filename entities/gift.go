package entities

import (
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Gift struct {
	gorm.Model
	GiftName   string   `gorm:"type:varchar(255); not null"`
	Total      uint     `gorm:"not null"`
	Remaining  uint     `gorm:"not null"`
	ProviderID uint     `json:"-"`
	Provider   Provider `json:"-" gorm:"foreignKey:ProviderID"`
}

func NewGift(dto dto.GiftDTO, providerID uint) Gift {
	return Gift{
		GiftName:   dto.GiftName,
		ProviderID: providerID,
		Total:      dto.Total,
		Remaining:  dto.Remaining,
	}
}
