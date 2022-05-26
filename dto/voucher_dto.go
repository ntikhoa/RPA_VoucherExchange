package dto

import "github.com/RPA_VoucherExchange/entities"

type VoucherDTO struct {
	Name               string              `json:"name" binding:"required"`
	Description        string              `json:"description" binding:"required"`
	TotalPriceMin      uint                `json:"total_price_min"`
	TotalPriceMax      uint                `json:"total_price_max"`
	ToTal              uint                `json:"total"`
	Remaining          uint                `json:"remaining"`
	Published          bool                `json:"published"`
	VoucherProductDTOs []VoucherProductDTO `json:"products" binding:"required"`
	Gift               string              `json:"gift" binding:"required"`
}

func (dto VoucherDTO) GetProductIDs() []uint {
	var productIDs []uint
	for _, p := range dto.VoucherProductDTOs {
		productIDs = append(productIDs, p.ProductID)
	}
	return productIDs
}

func (dto VoucherDTO) ToEntity(providerID uint) entities.Voucher {
	var voucherProducts []entities.VoucherProduct
	for _, x := range dto.VoucherProductDTOs {
		voucherProducts = append(voucherProducts, x.ToEntity())
	}

	return entities.Voucher{
		Name:            dto.Name,
		Description:     dto.Description,
		TotalPriceMin:   dto.TotalPriceMin,
		TotalPriceMax:   dto.TotalPriceMax,
		Total:           dto.ToTal,
		Remaining:       dto.Remaining,
		Published:       dto.Published,
		VoucherProducts: voucherProducts,
		Gift:            entities.Gift{GiftName: dto.Gift},
		ProviderID:      providerID,
	}
}
