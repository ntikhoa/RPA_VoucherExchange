package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductName string `gorm:"type:varchar(255); not null"`
}
