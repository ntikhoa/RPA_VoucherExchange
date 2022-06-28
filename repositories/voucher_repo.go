package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type VoucherRepo interface {
	Create(voucher entities.Voucher) error
	Update(voucher entities.Voucher) error
	FindByID(voucherID uint) (entities.Voucher, error)
	Delete(voucherID uint) error
	FindAllWithPage(providerID uint, page int, perPage int) ([]viewmodel.VoucherResponse, error)
	Count(providerID uint) (int64, error)
	Publish(voucherID uint, published bool) error
	FindVoucherExchange(providerID uint, productsName []string) ([]entities.Voucher, error)
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

func (r *voucherRepo) Update(voucher entities.Voucher) error {

	gift := entities.Gift{
		VoucherID: voucher.Model.ID,
	}
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&gift).Where(&gift).Update("gift_name", voucher.Gift.GiftName).Error; err != nil {
			return err
		}

		if err := tx.Model(&voucher).Association("Products").Replace(voucher.Products); err != nil {
			return err
		}

		if err := tx.Save(&voucher).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *voucherRepo) FindByID(voucherID uint) (entities.Voucher, error) {
	voucher := entities.Voucher{
		Model: gorm.Model{
			ID: voucherID,
		},
	}
	err := r.db.
		Preload("Products").
		Preload("Gift").
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

func (r *voucherRepo) Delete(voucherID uint) error {
	voucher := entities.Voucher{
		Model: gorm.Model{
			ID: voucherID,
		},
	}
	return r.db.Select("Gift", "Products").Delete(&voucher).Error
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

func (r *voucherRepo) FindVoucherExchange(providerID uint, productsName []string) ([]entities.Voucher, error) {
	var vouchers []entities.Voucher

	err := r.db.
		Preload("Gift").
		Preload("Products",
			"product_name IN (?) AND provider_id = ?",
			productsName,
			providerID).
		Find(&vouchers).
		Error
	return vouchers, err
}

//Find vouchers whose all products are in the receipt
//First find all vouchers id and then use gorm.preload to populate Produtcs and Gift
//To find all vouchers id, use the raw sql below
//SQL description:
//Join vouchers and products where product_name is in the receipt list
//Select and Group By voucher id and count(product_name)
//Do the same thing but without checking whether product_name is in the list
//Join the two corresponding tables with voucher_id and count_products
// func (r *voucherRepo) FindVoucherExchange(providerID uint) ([]entities.Voucher, error) {
// 	sql := `SELECT in_receipt.id FROM
// 	(
// 	SELECT v.id, COUNT(product_name) AS count_products
// 	FROM vouchers AS v
// 		INNER JOIN voucher_products AS vp ON vp.voucher_id = v.id AND provider_id = ?
// 		INNER JOIN products AS p ON vp.product_id = p.id AND p.product_name IN ('TN-MAM LUA MACH', 'HAU NUONG HASUBI200G')
// 	GROUP BY v.id
// 	) AS in_receipt
// 	INNER JOIN
// 	(
// 	SELECT v.id, COUNT(product_name) AS count_products
// 	FROM vouchers AS v
// 		INNER JOIN voucher_products AS vp ON vp.voucher_id = v.id AND provider_id = ?
// 		INNER JOIN products AS p ON vp.product_id = p.id
// 	GROUP BY v.id
// 	) AS all_vouchers
// 	ON in_receipt.id = all_vouchers.id AND in_receipt.count_products = all_vouchers.count_products`

// 	var vouchers []entities.Voucher
// 	var voucherIDs []uint
// 	err := r.db.
// 		Raw(sql, providerID, providerID).
// 		Scan(&voucherIDs).
// 		Error
// 	if err != nil {
// 		return vouchers, err
// 	}
// 	for _, value := range voucherIDs {
// 		log.Println(value)
// 	}

// 	err = r.db.
// 		Preload("Products").
// 		Preload("Gift").
// 		Where("id IN (?)", voucherIDs).
// 		Find(&vouchers).
// 		Error
// 	return vouchers, err
// }
