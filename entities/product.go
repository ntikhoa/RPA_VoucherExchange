package entities

import (
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string   `gorm:"type:varchar(255); not null; UNIQUE"`
	ProviderID  uint     `json:"-"`
	Provider    Provider `json:"-" gorm:"foreignKey:ProviderID"`
}

func NewProduct(dto dto.ProductDTO, providerID uint) Product {
	return Product{
		ProductName: dto.ProductName,
		ProviderID:  providerID,
	}
}
