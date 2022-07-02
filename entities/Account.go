package entities

import (
	"time"

	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username       string `gorm:"type:varchar(255); not null; UNIQUE"`
	HashedPassword string `gorm:"type:varchar(255); not null"`
	Name           string `gorm:"type:varchar(255)"`
	ProviderID     uint
	Provider       Provider `gorm:"foreignKey:ProviderID"`
	RoleID         uint
	Role           Role `gorm:"foreignKey:RoleID"`
	IssueAt        time.Time
}

func NewAccountAdmin(dto dto.RegisterDTO, hashedPassword string, role uint) Account {
	return Account{
		Username:       dto.Username,
		HashedPassword: hashedPassword,
		Name:           dto.Name,
		ProviderID:     dto.ProviderID,
		RoleID:         role,
		IssueAt:        time.Now(),
	}
}

func NewAccountSale(dto dto.RegisterSaleDTO, hashedPassword string, role uint, providerID uint) Account {
	return Account{
		Username:       dto.Username,
		HashedPassword: hashedPassword,
		Name:           dto.Name,
		ProviderID:     providerID,
		RoleID:         role,
		IssueAt:        time.Now(),
	}
}
