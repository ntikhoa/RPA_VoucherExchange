package dto

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type VoucherDTO struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	TotalPriceMin uint   `json:"total_price_min"`
	TotalPriceMax uint   `json:"total_price_max"`
	ToTal         uint   `json:"total"`
	Remaining     uint   `json:"remaining"`
	Published     bool   `json:"published"`
	ProductIDs    []uint `json:"products" binding:"required"`
	Gift          string `json:"gift" binding:"required"`
}

func (dto VoucherDTO) ToEntity(providerID uint) entities.Voucher {
	var products []entities.Product
	for _, x := range dto.ProductIDs {
		products = append(products, entities.Product{
			Model: gorm.Model{ID: x},
		})
	}

	return entities.Voucher{
		Name:          dto.Name,
		Description:   dto.Description,
		TotalPriceMin: dto.TotalPriceMin,
		TotalPriceMax: dto.TotalPriceMax,
		Total:         dto.ToTal,
		Remaining:     dto.Remaining,
		Published:     dto.Published,
		Products:      products,
		Gift: entities.Gift{
			GiftName:   dto.Gift,
			ProviderID: providerID,
		},
		ProviderID: providerID,
	}
}
