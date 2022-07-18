package entities

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	GiftName   string   `gorm:"type:varchar(255); not null"`
	ProviderID uint     `json:"-"`
	Provider   Provider `gorm:"foreignKey:ProviderID"`
}
