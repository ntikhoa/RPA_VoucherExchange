package dto

type VoucherDTO struct {
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	TotalPriceMin uint   `json:"total_price_min"`
	TotalPriceMax uint   `json:"total_price_max"`
	ToTal         uint   `json:"total"`
	Remaining     uint   `json:"remaining"`
	Published     bool   `json:"published"`
	ProductIDs    []uint `json:"products" binding:"required"`
	GiftID        uint   `json:"gift_id" binding:"required"`
}
