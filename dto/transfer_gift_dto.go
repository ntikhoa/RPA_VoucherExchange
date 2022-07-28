package dto

type TransferGiftDTO struct {
	GiftID   uint `form:"gift_id" json:"gift_id" binding:"required"`
	Quantity uint `form:"quantity" json:"quantity" binding:"required,gte=1"`
}

type CreateTransferGiftsDTO struct {
	TransferGiftDTO []TransferGiftDTO `json:"transfers" binding:"required"`
	AccountID       uint              `form:"account_id" json:"account_id" binding:"required"`
}
