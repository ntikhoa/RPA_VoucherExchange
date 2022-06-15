package entities

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Description string `gorm:"type:enum('ADMIN', 'SALE'); not null; UNIQUE"`
}
