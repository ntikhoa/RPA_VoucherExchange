package repositories

import (
	"time"

	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ReceiptRepo interface {
	Create(receipt entities.Receipt) error
	Count(providerID uint) (int64, error)
	FindAllWithPage(providerID uint) ([]entities.Receipt, error)
	FindByID(providerID uint, receiptID uint) (entities.Receipt, error)
	FindBetweenDates(productID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error)
	FindByIDWithoutJoin(providerID uint, receiptID uint) (entities.Receipt, error)
	UpdateCensorStatus(receiptID uint, statusID uint) error
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
		Model(&receipts).
		Preload("Voucher", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("provider_id = ?", providerID).Select("ID", "Name")
		}).
		Preload("Account", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("ID", "Name")
		}).
		Preload("Status").
		Joins("JOIN receipt_voucher ON receipts.id = receipt_voucher.receipt_id").
		Joins("JOIN vouchers ON vouchers.id = receipt_voucher.voucher_id AND vouchers.provider_id = ?", providerID).
		Find(&receipts).
		Error

	return receipts, err
}

func (r *receiptRepo) Count(providerID uint) (int64, error) {
	var count int64
	err := r.db.
		Model(&entities.Receipt{}).
		Joins("JOIN accounts ON accounts.id = receipts.account_id AND accounts.provider_id = ?", providerID).
		Count(&count).
		Error
	return count, err
}

func (r *receiptRepo) FindByID(providerID uint, receiptID uint) (entities.Receipt, error) {
	receipt := entities.Receipt{Model: gorm.Model{ID: receiptID}}
	err := r.db.
		Model(&receipt).
		Preload("Status").
		Preload("ReceiptItems").
		Preload("Voucher").
		Preload("Voucher.Products").
		Preload("Voucher.Gift").
		Preload("ReceiptImages").
		Preload("Customer").
		Preload("Account").
		Preload("Account.Role").
		Joins("JOIN receipt_voucher ON receipts.id = receipt_voucher.receipt_id").
		Joins("JOIN vouchers ON vouchers.id = receipt_voucher.voucher_id AND vouchers.provider_id = ?", providerID).
		First(&receipt).
		Error

	return receipt, err
}

func (r *receiptRepo) FindByIDWithoutJoin(providerID uint, receiptID uint) (entities.Receipt, error) {
	receipt := entities.Receipt{Model: gorm.Model{ID: receiptID}}
	err := r.db.
		Model(&receipt).
		Preload("Account").
		First(&receipt).
		Error

	return receipt, err
}

func (r *receiptRepo) FindBetweenDates(productID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error) {
	var receipts []entities.Receipt
	err := r.db.
		Model(&entities.Receipt{}).
		Where("DATE(created_at) BETWEEN DATE(?) AND DATE(?)", fromDate, toDate).
		Find(&receipts).
		Error
	return receipts, err

}

func (r *receiptRepo) UpdateCensorStatus(receiptID uint, statusID uint) error {
	return r.db.
		Model(&entities.Receipt{Model: gorm.Model{ID: receiptID}}).
		Update("status_id", statusID).
		Error
}
