package entities

import (
	"time"

	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username       string   `gorm:"type:varchar(255); not null; UNIQUE"`
	HashedPassword string   `json:"-" gorm:"type:varchar(255); not null"`
	Name           string   `gorm:"type:varchar(255)"`
	ProviderID     uint     `json:"-"`
	Provider       Provider `gorm:"foreignKey:ProviderID"`
	RoleID         uint     `json:"-"`
	Role           Role     `gorm:"foreignKey:RoleID"`
	IssueAt        time.Time
}

func NewAccount(dto dto.RegisterDTO, hashedPassword string, role uint) Account {
	return Account{
		Username:       dto.Username,
		HashedPassword: hashedPassword,
		Name:           dto.Name,
		ProviderID:     dto.ProviderID,
		RoleID:         role,
		IssueAt:        time.Now(),
	}
}
