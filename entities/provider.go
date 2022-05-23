package entities

import "gorm.io/gorm"

type Provider struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(255); not null; UNIQUE"`
	Products []Product `gorm:"foreignKey:ProviderID"`
	Gifts    []Gift    `gorm:"foreignKey:ProviderID"`
}
