package dto

import "github.com/RPA_VoucherExchange/entities"

type VoucherDTO struct {
	Name              string              `json:"name" binding:"required"`
	Description       string              `json:"description" binding:"required"`
	TotalPriceMin     uint                `json:"total_price_min"`
	TotalPriceMax     uint                `json:"total_price_max"`
	ToTal             uint                `json:"total"`
	Remaining         uint                `json:"remaining"`
	Published         bool                `json:"published"`
	VoucherProductDTO []VoucherProductDTO `json:"products" binding:"required"`
	Gift              []string            `json:"gifts" binding:"required"`
}

func (dto VoucherDTO) GetVoucherProducts() []uint {
	var productIDs []uint
	for _, p := range dto.VoucherProductDTO {
		productIDs = append(productIDs, p.ProductID)
	}
	return productIDs
}

func (dto VoucherDTO) ToEntity(providerID uint) entities.Voucher {
	var voucherProduct []entities.VoucherProduct
	for _, x := range dto.VoucherProductDTO {
		voucherProduct = append(voucherProduct, x.ToEntity())
	}
	var gifts []entities.Gift
	for _, x := range dto.Gift {
		gifts = append(gifts, entities.Gift{
			GiftName:   x,
			ProviderID: providerID,
		})
	}

	return entities.Voucher{
		Name:           dto.Name,
		Description:    dto.Description,
		TotalPriceMin:  dto.TotalPriceMin,
		TotalPriceMax:  dto.TotalPriceMax,
		Total:          dto.ToTal,
		Remaining:      dto.Remaining,
		Published:      dto.Published,
		VoucherProduct: voucherProduct,
		Gift:           gifts,
		ProviderID:     providerID,
	}
}
