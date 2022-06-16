package entities

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Username       string `gorm:"type:varchar(255); not null; UNIQUE"`
	HashedPassword string `gorm:"type:varchar(255); not null"`
	Name           string `gorm:"type:varchar(255)"`
	ProviderID     uint
	Provider       Provider `gorm:"foreignKey:ProviderID"`
	RoleID         uint
	Role           Role `gorm:"foreignKey:RoleID"`
	IssueAt        time.Time
}
