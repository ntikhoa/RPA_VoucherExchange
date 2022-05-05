package utils

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) {
	db.Create(&entities.Provider{
		Name: "CONG TY TNHH CONG DONG VIET",
	})
}
