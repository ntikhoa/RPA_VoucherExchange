package entities

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Username   string `gorm:"type:varchar(255); not null; UNIQUE"`
	Password   string `gorm:"type:varchar(255); not null"`
	Name       string `gorm:"type:varchar(255)"`
	ProviderID int
	Provider   Provider `gorm:"foreignKey:ProviderID"`
}
