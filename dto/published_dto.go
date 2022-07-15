package dto

type PublishedDTO struct {
	VoucherID uint `form:"voucher_id" json:"voucher_id" binding:"required"`
	Published bool `form:"published" json:"published"`
}
