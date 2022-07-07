package viewmodel

import (
	"time"

	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type AccountResponse struct {
	ID        uint
	Username  string
	Name      string
	RoleID    uint
	Role      entities.Role `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
