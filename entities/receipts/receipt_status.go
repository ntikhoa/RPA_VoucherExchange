package receipts

import (
	"gorm.io/gorm"
)

type ReceiptStatus struct {
	gorm.Model
	Description string `gorm:"type:enum('PENDING', 'COMPLETE')"`
}
