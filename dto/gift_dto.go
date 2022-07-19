package dto

type GiftDTO struct {
	GiftName  string `form:"gift_name" json:"gift_name" binding:"required"`
	Total     uint   `form:"total" json:"total"`
	Remaining uint   `form:"remaining" json:"remaining"`
}
