package dto

import "github.com/RPA_VoucherExchange/entities"

type RegisterDTO struct {
	Username          string `form:"username" binding:"required"`
	Password          string `form:"password" binding:"required,gte=6"`
	ConfirmedPassword string `form:"confirmed_password" binding:"required,gte=6"`
	Name              string `form:"name"`
	ProviderID        uint   `form:"provider_id" binding:"required"`
}

func (dto RegisterDTO) ToEmployeeEntity() entities.Employee {
	return entities.Employee{
		Username:   dto.Username,
		Name:       dto.Name,
		ProviderID: dto.ProviderID,
	}
}
