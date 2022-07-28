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

func NewTransfer(dto dto.TransferGiftDTO, accountID uint) Transfer {
	return Transfer{
		AccountID: accountID,
		Quantity:  dto.Quantity,
		GiftID:    dto.GiftID,
	}
}

func NewSliceTransfer(dto dto.CreateTransferGiftsDTO) []Transfer {
	var transfer []Transfer
	for _, value := range dto.TransferGiftDTO {
		transfer = append(transfer, NewTransfer(value, dto.AccountID))
	}
	return transfer
}
