package entities

import (
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `gorm:"type:varchar(255); not null; UNIQUE; index"`
	ProviderID  uint   `json:"-"`
}

func NewProduct(dto dto.ProductDTO, providerID uint) Product {
	return Product{
		ProductName: dto.ProductName,
		ProviderID:  providerID,
	}
}
