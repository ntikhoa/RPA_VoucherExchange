package viewmodel

import (
	"time"

	"gorm.io/gorm"
)

type VoucherResponse struct {
	ID          uint
	Name        string
	Description string
	Total       uint
	Remaining   uint
	Published   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
