package viewmodel

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type AccountResponse struct {
	gorm.Model
	Username   string
	Name       string
	RoleID     uint          `json:"-"`
	Role       entities.Role `gorm:"foreignKey:RoleID"`
	ProviderID uint          `json:"-"`
}

func NewAccountResponse(entity entities.Account) AccountResponse {
	return AccountResponse{
		Model:      entity.Model,
		Username:   entity.Username,
		Name:       entity.Name,
		Role:       entity.Role,
		ProviderID: entity.ProviderID,
	}
}

type AccountProfile struct {
	gorm.Model
	Username string
	Name     string
	Provider entities.Provider
}

func NewAccountProfileRes(entity entities.Account) AccountProfile {
	return AccountProfile{
		Model:    entity.Model,
		Username: entity.Username,
		Name:     entity.Name,
		Provider: entity.Provider,
	}
}
