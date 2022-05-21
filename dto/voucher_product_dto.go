package dto

type VoucherProductDTO struct {
	ProductID   uint `json:"product_id" binding:"required"`
	QuantityMin uint `json:"quantity_min" binding:"required"`
	QuantityMax uint `json:"quantity_max" binding:"required"`
}
