package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ReceiptRepo interface {
	Create(receipt entities.Receipt) error
	Count(providerID uint) (int64, error)
	FindAllWithPage(providerID uint) ([]entities.Receipt, error)
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

func (r *receiptRepo) FindAllWithPage(providerID uint) ([]entities.Receipt, error) {
	var receipts []entities.Receipt

	err := r.db.
		Model(&entities.Receipt{}).
		Preload("Voucher", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("provider_id = ?", providerID).Select("ID", "Name")
		}).
		Preload("Account", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("ID", "Username")
		}).
		Preload("Status").
		Find(&receipts).
		Error

	return receipts, err
}

func (r *receiptRepo) Count(providerID uint) (int64, error) {
	var count int64
	err := r.db.
		Model(&entities.Receipt{}).
		Preload("Account", "provider_id = ?", providerID).
		Count(&count).
		Error
	return count, err
}
