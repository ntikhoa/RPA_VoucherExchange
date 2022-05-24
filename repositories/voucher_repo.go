package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	Create(voucher entities.Voucher) error
	FindByID(voucherID uint) (entities.Voucher, error)
	FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Voucher, error)
	Count(providerID uint) (int64, error)
}

type voucherRepo struct {
	db *gorm.DB
}

func NewVoucherRepo(db *gorm.DB) VoucherRepo {
	return &voucherRepo{
		db: db,
	}
}

func (r *voucherRepo) Create(voucher entities.Voucher) error {
	return r.db.Create(&voucher).Error
}

func (r *voucherRepo) FindByID(voucherID uint) (entities.Voucher, error) {
	voucher := entities.Voucher{
		Model: gorm.Model{
			ID: voucherID,
		},
	}
	err := r.db.
		Preload("VoucherProducts").
		Preload("Gifts").
		First(&voucher).
		Error
	return voucher, err
}

func (r *voucherRepo) Count(providerID uint) (int64, error) {
	var count int64
	err := r.db.
		Model(&entities.Voucher{ProviderID: providerID}).
		Count(&count).
		Error
	return count, err
}

func (r *voucherRepo) FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Voucher, error) {
	var vouchers []entities.Voucher
	err := r.db.
		Where(&entities.Voucher{
			ProviderID: providerID,
		}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&vouchers).
		Error
	return vouchers, err
}
