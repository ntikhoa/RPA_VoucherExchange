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
	Total         uint      `gorm:"not null"`
	Remaining     uint      `gorm:"not null"`
	Published     bool      `gorm:"not null"`
	Products      []Product `gorm:"many2many:voucher_products;constraint:OnDelete:CASCADE"`
	Gift          Gift      `gorm:"foreignKey:VoucherID;constraint:OnDelete:CASCADE"`
	ProviderID    uint      `json:"-"`
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
		Total:         dto.ToTal,
		Remaining:     dto.Remaining,
		Published:     dto.Published,
		Products:      products,
		Gift: Gift{
			GiftName:   dto.Gift,
			ProviderID: providerID,
		},
		ProviderID: providerID,
	}
}
