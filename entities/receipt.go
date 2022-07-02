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

func NewReceipt(dto dto.ExchangeVoucherDTO, filesNames []string, accountID uint) Receipt {

	receiptItems := NewReceiptItems(dto.ViewExchangeVoucherDTO)
	receiptImages := NewReceiptImages(filesNames)

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
