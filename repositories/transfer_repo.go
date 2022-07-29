package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type TransferRepo interface {
	CreateTransfers(transfer []entities.Transfer) error
	GetTransfersByAccount(accountID uint, providerID uint) ([]entities.Transfer, error)
	DeleteTransfers(accountID uint) error
}

type transferRepo struct {
	db *gorm.DB
}

func NewTransferRepo(db *gorm.DB) TransferRepo {
	return &transferRepo{db: db}
}

func (r *transferRepo) CreateTransfers(transfer []entities.Transfer) error {
	return r.db.Save(&transfer).Error
}

func (r *transferRepo) GetTransfersByAccount(accountID uint, providerID uint) ([]entities.Transfer, error) {
	var transfers []entities.Transfer

	err := r.db.Model(&transfers).
		Preload("Gift", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("ID", "GiftName")
		}).
		Joins("JOIN accounts ON transfers.account_id = ? AND accounts.id = transfers.account_id AND accounts.provider_id = ?", accountID, providerID).
		Find(&transfers).
		Error

	return transfers, err
}

func (r *transferRepo) DeleteTransfers(accountID uint) error {
	return r.db.
		Where("account_id = ?", accountID).
		Delete(&entities.Transfer{}).
		Error
}
