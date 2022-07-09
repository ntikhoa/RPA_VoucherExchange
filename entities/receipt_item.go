package entities

import (
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type ReceiptItem struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255); not null"`
	Amount    uint   `gorm:"not null"`
	ReceiptID uint   `json:"-"`
}

func NewReceiptItems(dto dto.ViewExchangeVoucherDTO) []ReceiptItem {
	var receiptItems []ReceiptItem
	for i := range dto.Products {
		receiptItems = append(receiptItems, ReceiptItem{
			Name:   dto.Products[i],
			Amount: dto.Prices[i],
		})
	}
	return receiptItems
}
