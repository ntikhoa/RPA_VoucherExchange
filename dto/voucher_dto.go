package dto

type VoucherDTO struct {
	Name              string              `json:"name" binding:"required"`
	Description       string              `json:"description" binding:"required"`
	TotalPriceMin     int                 `json:"total_price_min" binding:"required"`
	TotalPriceMax     int                 `json:"total_price_max" binding:"required"`
	VoucherProductDTO []VoucherProductDTO `json:"products" binding:"required"`
	Gift              []string            `json:"gifts" binding:"required"`
}

func (dto *VoucherDTO) GetVoucherProducts() []uint {
	var productIDs []uint
	for _, p := range dto.VoucherProductDTO {
		productIDs = append(productIDs, p.ProductID)
	}
	return productIDs
}
