package viewmodel

import (
	"gorm.io/gorm"
)

type VoucherResponse struct {
	gorm.Model
	Name        string
	Description string
	Published   bool
}
