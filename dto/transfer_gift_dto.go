package dto

type TransferGiftDTO struct {
	AccountID uint `form:"account_id" json:"account_id" binding:"required"`
	GiftID    uint `form:"gift_id" json:"gift_id" binding:"required"`
	Quantity  uint `form:"quantity" json:"quantity" binding:"required,gte=1"`
}
