package viewmodel

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type TransferResponse struct {
	gorm.Model
	Quantity uint
	GiftID   uint
	GiftName string
}

func NewTransferResponse(entity entities.Transfer) TransferResponse {
	return TransferResponse{
		Model:    entity.Model,
		Quantity: entity.Quantity,
		GiftID:   entity.GiftID,
		GiftName: entity.Gift.GiftName,
	}
}

func NewSliceTransferResponse(entities []entities.Transfer) []TransferResponse {
	var res []TransferResponse
	for _, entity := range entities {
		res = append(res, NewTransferResponse(entity))
	}
	return res
}
