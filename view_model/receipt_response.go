package viewmodel

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ReceiptListRes struct {
	gorm.Model
	TransactionID string
	Status        entities.CensorStatus
	Voucher       string
	Account       string
}

func NewReceiptListRes(entity entities.Receipt) ReceiptListRes {
	return ReceiptListRes{
		Model:         entity.Model,
		TransactionID: entity.TransactionID,
		Status:        entity.Status,
		Voucher:       entity.Voucher[0].Name,
		Account:       entity.Account.Username,
	}
}

func NewSliceReceiptListRes(entities []entities.Receipt) []ReceiptListRes {
	var results []ReceiptListRes
	for _, entity := range entities {
		results = append(results, NewReceiptListRes(entity))
	}
	return results
}
