package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ReceiptRepo interface {
	Create(receipt entities.Receipt) error
}

type receiptRepo struct {
	db *gorm.DB
}

func NewReceiptRepo(db *gorm.DB) ReceiptRepo {
	return &receiptRepo{
		db: db,
	}
}

func (r *receiptRepo) Create(receipt entities.Receipt) error {
	return r.db.Create(&receipt).Error
}
