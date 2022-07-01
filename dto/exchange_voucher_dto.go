package dto

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ViewExchangeVoucherDTO struct {
	Products []string
	Prices   []uint
}

type ExchangeVoucherDTO struct {
	ViewExchangeVoucherDTO ViewExchangeVoucherDTO
	TransactionID          string
	VoucherID              uint
	CustomerName           string
	CustomerPhone          string
}

func (dto *ExchangeVoucherDTO) ToEntitiy(filesName []string, accountID uint) entities.Receipt {
	var receiptItems []entities.ReceiptItem
	for i := range dto.ViewExchangeVoucherDTO.Products {
		receiptItems = append(receiptItems, entities.ReceiptItem{
			Name:   dto.ViewExchangeVoucherDTO.Products[i],
			Amount: dto.ViewExchangeVoucherDTO.Prices[i],
		})
	}

	var receiptImages []entities.ReceiptImage
	for _, fileName := range filesName {
		receiptImages = append(receiptImages, entities.ReceiptImage{
			Url: fileName,
		})
	}

	voucher := []entities.Voucher{{Model: gorm.Model{ID: dto.VoucherID}}}

	return entities.Receipt{
		StatusID:      constants.STATUS_PENDING,
		TransactionID: dto.TransactionID,
		ReceiptItem:   receiptItems,
		Voucher:       voucher,
		ReceiptImage:  receiptImages,
		Customer: entities.Customer{
			Name:  dto.CustomerName,
			Phone: dto.CustomerPhone,
		},
		AccountID: accountID,
	}
}
