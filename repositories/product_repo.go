package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(product entities.Product) error
	Update(product entities.Product) error
	DeleteByID(productID uint) error
	DeleteByIDs(productIDs []uint) error
	FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Product, error)
	FindByID(productID uint) (entities.Product, error)
	FindByIDs(productIDs []uint) ([]entities.Product, error)
	Search(query string, providerID uint) ([]entities.Product, error)
	Count(providerID uint) (int64, error)
	CheckExistence(providerID uint, productIDs []uint) ([]uint, error)
	GetAll(providerID uint) ([]entities.Product, error)
	GetExchangeProductsNames(providerID uint) ([]string, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (repo *productRepo) Create(product entities.Product) error {
	return repo.db.Create(&product).Error
}

func (repo *productRepo) Update(product entities.Product) error {
	return repo.db.Omit("created_at").Save(&product).Error
}

func (repo *productRepo) DeleteByID(productID uint) error {
	return repo.db.
		Unscoped().
		Delete(&entities.Product{}, productID).
		Error
}

func (repo *productRepo) FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Product, error) {
	products := []entities.Product{}
	tx := repo.db.
		Where(&entities.Product{ProviderID: providerID}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&products)
	return products, tx.Error
}

func (repo *productRepo) FindByID(productID uint) (entities.Product, error) {
	product := entities.Product{
		Model: gorm.Model{
			ID: productID,
		},
	}
	err := repo.db.First(&product).Error

	return product, err
}

func (repo *productRepo) Search(query string, providerID uint) ([]entities.Product, error) {
	products := []entities.Product{}
	err := repo.db.
		Where("provider_id = ? AND product_name LIKE ?", providerID, "%"+query+"%").
		Find(&products).
		Error
	return products, err
}

func (repo *productRepo) Count(providerID uint) (int64, error) {
	var count int64
	err := repo.db.
		Model(&entities.Product{ProviderID: providerID}).
		Count(&count).
		Error

	return count, err
}

type productID struct {
	ID uint
}

func (repo *productRepo) CheckExistence(providerID uint, productIDs []uint) ([]uint, error) {
	var IDs []productID
	tx := repo.
		db.
		Model(&entities.Product{}).
		Where("id IN (?) AND provider_id = ?", productIDs, providerID).
		Find(&IDs)

	ids := []uint{}

	for _, id := range IDs {
		ids = append(ids, id.ID)
	}
	return ids, tx.Error
}

func (repo *productRepo) GetAll(providerID uint) ([]entities.Product, error) {
	var products []entities.Product

	err := repo.
		db.
		Where(&entities.Product{ProviderID: providerID}).
		Order("product_name").
		Find(&products).
		Error

	return products, err
}

func (repo *productRepo) FindByIDs(productIDs []uint) ([]entities.Product, error) {
	var products []entities.Product
	err := repo.db.Find(&products, productIDs).Error
	return products, err
}

func (repo *productRepo) DeleteByIDs(productIDs []uint) error {
	var products []entities.Product
	return repo.db.
		Unscoped().
		Delete(&products, productIDs).
		Error
}

func (repo *productRepo) GetExchangeProductsNames(providerID uint) ([]string, error) {
	var products []string
	err := repo.db.
		Model(&entities.Product{}).
		Select("product_name").
		Joins("JOIN voucher_products ON products.id = voucher_products.product_id AND products.provider_id = ?", providerID).
		Joins("JOIN vouchers ON vouchers.id = voucher_products.voucher_id AND vouchers.published = 1").
		Distinct().
		Find(&products).
		Error

	return products, err
}
