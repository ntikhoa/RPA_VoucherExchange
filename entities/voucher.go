package entities

import (
	"github.com/RPA_VoucherExchange/dto"
	"gorm.io/gorm"
)

type Voucher struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(255)"`
	Description   string    `gorm:"type:varchar(255)"`
	TotalPriceMin uint      `gorm:"not null"`
	TotalPriceMax uint      `gorm:"not null"`
	Published     bool      `gorm:"not null"`
	Products      []Product `gorm:"many2many:voucher_products;constraint:OnDelete:CASCADE"`
	GiftID        uint      `json:"-"`
	Gift          Gift      `gorm:"foreignKey:GiftID"`
	ProviderID    uint      `json:"-"`
	Provider      Provider  `json:"-" gorm:"foreignKey:ProviderID"`
}

func NewVoucher(dto dto.VoucherDTO, providerID uint) Voucher {
	var products []Product
	for _, x := range dto.ProductIDs {
		products = append(products, Product{
			Model: gorm.Model{ID: x},
		})
	}

	return Voucher{
		Name:          dto.Name,
		Description:   dto.Description,
		TotalPriceMin: dto.TotalPriceMin,
		TotalPriceMax: dto.TotalPriceMax,
		Published:     dto.Published,
		Products:      products,
		GiftID:        dto.GiftID,
		ProviderID:    providerID,
	}
}
