package viewmodel

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type AccountResponse struct {
	gorm.Model
	Username string
	Name     string
	RoleID   uint          `json:"-"`
	Role     entities.Role `gorm:"foreignKey:RoleID"`
}

func NewAccountResponse(entity entities.Account) AccountResponse {
	return AccountResponse{
		Model:    entity.Model,
		Username: entity.Username,
		Name:     entity.Name,
		Role:     entity.Role,
	}
}
