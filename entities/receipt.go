package entities

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	TransactionID string `gorm:"type:varchar(10); not null; index; UNIQUE"`
	StatusID      uint
	Status        CensorStatus   `gorm:"foreignKey:StatusID"`
	ReceiptItem   []ReceiptItem  `gorm:"foreignKey:ReceiptID"`
	Voucher       []Voucher      `gorm:"many2many:receipt_voucher"`
	ReceiptImage  []ReceiptImage `gorm:"foreignKey:ReceiptID"`
	Customer      Customer       `gorm:"foreignKey:ReceiptID"`
	AccountID     uint
	Account       Account `gorm:"foreignKey:AccountID"`
}

func NewReceipt(dto dto.ExchangeVoucherDTO, filesName []string, accountID uint) Receipt {
	var receiptItems []ReceiptItem
	for i := range dto.ViewExchangeVoucherDTO.Products {
		receiptItems = append(receiptItems, ReceiptItem{
			Name:   dto.ViewExchangeVoucherDTO.Products[i],
			Amount: dto.ViewExchangeVoucherDTO.Prices[i],
		})
	}

	var receiptImages []ReceiptImage
	for _, fileName := range filesName {
		receiptImages = append(receiptImages, ReceiptImage{
			Url: fileName,
		})
	}

	voucher := []Voucher{{Model: gorm.Model{ID: dto.VoucherID}}}

	return Receipt{
		StatusID:      constants.STATUS_PENDING,
		TransactionID: dto.TransactionID,
		ReceiptItem:   receiptItems,
		Voucher:       voucher,
		ReceiptImage:  receiptImages,
		Customer: Customer{
			Name:  dto.CustomerName,
			Phone: dto.CustomerPhone,
		},
		AccountID: accountID,
	}
}
