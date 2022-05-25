package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	Create(voucher entities.Voucher) error
	FindByID(voucherID uint) (entities.Voucher, error)
	Delete(voucherID uint) error
	FindAllWithPage(providerID uint, page int, perPage int) ([]viewmodel.VoucherResponse, error)
	Count(providerID uint) (int64, error)
	Publish(voucherID uint, published bool) error
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

func (r *voucherRepo) FindAllWithPage(providerID uint, page int, perPage int) ([]viewmodel.VoucherResponse, error) {
	var vouchersRes []viewmodel.VoucherResponse
	err := r.db.
		Model(&entities.Voucher{}).
		Where(&entities.Voucher{
			ProviderID: providerID,
		}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&vouchersRes).
		Error
	return vouchersRes, err
}

//transaction delete voucher product, gift and voucher with the voucher_id field
func (r *voucherRepo) Delete(voucherID uint) error {
	voucher := entities.Voucher{
		Model: gorm.Model{
			ID: voucherID,
		},
	}
	voucherProduct := entities.VoucherProduct{
		VoucherID: voucherID,
	}
	gift := entities.Gift{
		VoucherID: &voucherID,
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&voucherProduct).
			Delete([]entities.VoucherProduct{}).
			Error; err != nil {
			return err
		}

		if err := tx.Where(&gift).Delete([]entities.Gift{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&voucher).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *voucherRepo) Publish(voucherID uint, published bool) error {
	voucher := entities.Voucher{
		Model: gorm.Model{
			ID: voucherID,
		},
	}

	return r.db.
		Model(&voucher).
		Update("published", published).
		Error
}
