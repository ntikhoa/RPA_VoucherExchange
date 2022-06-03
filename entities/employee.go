package entities

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Username       string `gorm:"type:varchar(255); not null; UNIQUE"`
	HashedPassword string `gorm:"type:varchar(255); not null"`
	Name           string `gorm:"type:varchar(255)"`
	ProviderID     uint
	Provider       Provider `gorm:"foreignKey:ProviderID"`
	IssueAt        time.Time
}
