package dto

import "github.com/RPA_VoucherExchange/entities"

type ProductDTO struct {
	ProductName string `form:"product_name" json:"product_name" binding:"required"`
}

func (dto ProductDTO) ToEntity(providerID uint) entities.Product {
	return entities.Product{
		ProductName: dto.ProductName,
		ProviderID:  providerID,
	}
}
