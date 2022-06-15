package dto

type PayLoad struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}
