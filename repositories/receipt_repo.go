package repositories

import (
	"log"
	"time"

	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ReceiptRepo interface {
	Create(receipt entities.Receipt) error
	Count(providerID uint) (int64, error)
	CountByAccount(providerID uint, accountID uint) (int64, error)
	FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Receipt, error)
	FindAllWithAccountWithPage(providerID uint, accountID uint, page int, perPage int) ([]entities.Receipt, error)
	FindByID(providerID uint, receiptID uint) (entities.Receipt, error)
	FindBetweenDates(providerID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error)
	FindBetweenDatesWithAccount(providerID uint, accountID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error)
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

func (r *receiptRepo) FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Receipt, error) {
	var receipts []entities.Receipt

	err := r.db.
		Model(&receipts).
		Preload("Voucher", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Account", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Status").
		Joins("JOIN receipt_voucher ON receipts.id = receipt_voucher.receipt_id").
		Joins("JOIN vouchers ON vouchers.id = receipt_voucher.voucher_id AND vouchers.provider_id = ?", providerID).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&receipts).
		Error

	return receipts, err
}

func (r *receiptRepo) FindAllWithAccountWithPage(providerID uint, accountID uint, page int, perPage int) ([]entities.Receipt, error) {
	var receipts []entities.Receipt

	log.Println(providerID, accountID)

	err := r.db.
		Model(&receipts).
		Preload("Voucher", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Voucher.Gift", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "GiftName")
		}).
		Preload("Status").
		Joins("JOIN accounts ON accounts.id IN (?) AND accounts.provider_id = ? AND accounts.id = receipts.account_id", accountID, providerID).
		Joins("JOIN receipt_voucher ON receipts.id = receipt_voucher.receipt_id").
		Joins("JOIN vouchers ON vouchers.id = receipt_voucher.voucher_id AND vouchers.provider_id = ?", providerID).
		Limit(perPage).
		Offset((page - 1) * perPage).
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

func (r *receiptRepo) CountByAccount(providerID uint, accountID uint) (int64, error) {
	log.Println(providerID, accountID)

	var count int64
	err := r.db.
		Model(&entities.Receipt{}).
		Joins("JOIN accounts ON accounts.id = ? AND accounts.provider_id = ? AND accounts.id = receipts.account_id", accountID, providerID).
		Count(&count).
		Error

	log.Println(count)
	return count, err
}

func (r *receiptRepo) FindByID(providerID uint, receiptID uint) (entities.Receipt, error) {
	receipt := entities.Receipt{Model: gorm.Model{ID: receiptID}}
	err := r.db.
		Model(&receipt).
		Unscoped().
		Preload("Status").
		Preload("ReceiptItems").
		Preload("Voucher", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("Voucher.Products").
		Preload("Voucher.Gift").
		Preload("ReceiptImages").
		Preload("Customer").
		Preload("Account").
		Preload("Account.Role").
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

func (r *receiptRepo) FindBetweenDates(providerID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error) {
	var receipts []entities.Receipt
	err := r.db.
		Model(&entities.Receipt{}).
		Preload("Voucher", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Account", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Status").
		Joins("JOIN receipt_voucher ON receipts.id = receipt_voucher.receipt_id AND DATE(receipts.created_at) BETWEEN DATE(?) AND DATE(?)", fromDate, toDate).
		Joins("JOIN vouchers ON vouchers.id = receipt_voucher.voucher_id AND vouchers.provider_id = ?", providerID).
		Find(&receipts).
		Error
	return receipts, err

}

func (r *receiptRepo) FindBetweenDatesWithAccount(providerID uint, accountID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error) {
	var receipts []entities.Receipt
	err := r.db.
		Model(&entities.Receipt{}).
		Preload("Voucher", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Account", func(tx *gorm.DB) *gorm.DB {
			return tx.Unscoped().Select("ID", "Name")
		}).
		Preload("Status").
		Joins("JOIN accounts ON accounts.id IN (?) AND accounts.provider_id = ? AND accounts.id = receipts.account_id", accountID, providerID).
		Joins("JOIN receipt_voucher ON receipts.id = receipt_voucher.receipt_id AND DATE(receipts.created_at) BETWEEN DATE(?) AND DATE(?)", fromDate, toDate).
		Joins("JOIN vouchers ON vouchers.id = receipt_voucher.voucher_id AND vouchers.provider_id = ?", providerID).
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
