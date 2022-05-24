package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	Create(voucher entities.Voucher) error
	FindByID(voucherID uint) (entities.Voucher, error)
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
