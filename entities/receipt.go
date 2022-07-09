package entities

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	TransactionID string         `gorm:"type:varchar(10); not null; index"`
	StatusID      uint           `json:"-"`
	Status        CensorStatus   `gorm:"foreignKey:StatusID"`
	ReceiptItems  []ReceiptItem  `gorm:"foreignKey:ReceiptID"`
	Voucher       []Voucher      `gorm:"many2many:receipt_voucher"`
	ReceiptImages []ReceiptImage `gorm:"foreignKey:ReceiptID"`
	Customer      Customer       `gorm:"foreignKey:ReceiptID"`
	AccountID     uint           `json:"-"`
	Account       Account        `json:"-" gorm:"foreignKey:AccountID"`
}

func NewReceipt(dto dto.ExchangeVoucherDTO, filesNames []string, accountID uint) Receipt {

	receiptItems := NewReceiptItems(dto.ViewExchangeVoucherDTO)
	receiptImages := NewReceiptImages(filesNames)

	voucher := []Voucher{{Model: gorm.Model{ID: dto.VoucherID}}}

	return Receipt{
		StatusID:      constants.STATUS_PENDING,
		TransactionID: dto.TransactionID,
		ReceiptItems:  receiptItems,
		Voucher:       voucher,
		ReceiptImages: receiptImages,
		Customer: Customer{
			Name:  dto.CustomerName,
			Phone: dto.CustomerPhone,
		},
		AccountID: accountID,
	}
}
