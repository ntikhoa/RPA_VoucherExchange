package entities

import "gorm.io/gorm"

type Gift struct {
	gorm.Model
	GiftName   string   `gorm:"type:varchar(255); not null"`
	Total      uint     `gorm:"not null"`
	Remaining  uint     `gorm:"not null"`
	ProviderID uint     `json:"-"`
	Provider   Provider `gorm:"foreignKey:ProviderID"`
}
