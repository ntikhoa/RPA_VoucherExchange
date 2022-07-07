package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type TransactionRepo interface {
	FindAll(providerID uint) ([]entities.Receipt, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) TransactionRepo {
	return &transactionRepo{db: db}
}

func (r *transactionRepo) FindAll(providerID uint) ([]entities.Receipt, error) {
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
