package entities

import (
	"gorm.io/gorm"
)

type CensorStatus struct {
	gorm.Model
	Description string `gorm:"type:enum('PENDING', 'APPROVED', 'REJECTED'); not null; UNIQUE"`
}
