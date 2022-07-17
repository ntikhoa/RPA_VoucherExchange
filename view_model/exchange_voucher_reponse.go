package viewmodel

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ExchangeVoucherResponse struct {
	gorm.Model
	Name        string
	Description string
	GiftName    string
}

func NewExchangeVoucherResponse(entity entities.Voucher) ExchangeVoucherResponse {
	return ExchangeVoucherResponse{
		Model:       entity.Model,
		Name:        entity.Name,
		Description: entity.Description,
		GiftName:    entity.Gift.GiftName,
	}
}

func NewExchangeVoucherListResponse(entities []entities.Voucher) []ExchangeVoucherResponse {
	var res []ExchangeVoucherResponse
	for _, value := range entities {
		res = append(res, NewExchangeVoucherResponse(value))
	}
	return res
}
