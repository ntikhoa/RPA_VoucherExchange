package dto

type GiftDTO struct {
	GiftName  string `form:"gift_name" json:"gift_name" binding:"required"`
	VoucherID uint   `form:"voucher_id" json:"voucher_id" binding:"required"`
	Total     uint   `form:"total" json:"total" binding:"required"`
	Remaining uint   `form:"remaining" json:"remaining" binding:"required"`
}
