package viewmodel

import (
	"gorm.io/gorm"
)

type VoucherResponse struct {
	gorm.Model
	Name        string
	Description string
	Total       uint
	Remaining   uint
	Published   bool
}
