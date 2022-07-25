package entities

import (
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Transfer struct {
	gorm.Model
	AccountID uint    `json:"-"`
	Account   Account `json:"-" gorm:"foreignKey:AccountID"`
	Quantity  uint    `gorm:"not null"`
	GiftID    uint    `json:"-"`
	Gift      Gift    `gorm:"foreignKey:GiftID"`
}

func NewTransfer(dto dto.TransferGiftDTO) Transfer {
	return Transfer{
		AccountID: dto.AccountID,
		Quantity:  dto.Quantity,
		GiftID:    dto.GiftID,
	}
}
