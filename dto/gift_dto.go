package dto

import "github.com/RPA_VoucherExchange/entities"

type GiftDTO struct {
	GiftName string `form:"gift_name" json:"gift_name" binding:"required"`
}

func (dto GiftDTO) ToEntity() entities.Gift {
	return entities.Gift{
		GiftName: dto.GiftName,
	}
}
