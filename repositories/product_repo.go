package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ProductRepo interface {
	CreateProduct(product entities.Product) error
	UpdateProduct(product entities.Product) error
	DeleteProductByID(productID uint) error
	FindAllProductWithPage(providerID uint, page int, perPage int) ([]entities.Product, error)
	FindProductByID(productID uint) (entities.Product, error)
	GetProductCount(providerID uint) (int64, error)

	CheckProductsExist(productIDs []uint) ([]uint, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (repo *productRepo) CreateProduct(product entities.Product) error {
	return repo.db.Create(&product).Error
}

func (repo *productRepo) UpdateProduct(product entities.Product) error {
	return repo.db.Save(&product).Error
}

func (repo *productRepo) DeleteProductByID(productID uint) error {
	return repo.db.
		Unscoped().
		Delete(&entities.Product{}, productID).
		Error
}

func (repo *productRepo) FindAllProductWithPage(providerID uint, page int, perPage int) ([]entities.Product, error) {
	products := []entities.Product{}
	tx := repo.db.
		Where(&entities.Product{ProviderID: providerID}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&products)
	return products, tx.Error
}

func (repo *productRepo) FindProductByID(productID uint) (entities.Product, error) {
	product := entities.Product{
		Model: gorm.Model{
			ID: productID,
		},
	}
	err := repo.db.First(&product).Error

	return product, err
}

func (repo *productRepo) GetProductCount(providerID uint) (int64, error) {
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

func (repo *productRepo) CheckProductsExist(productIDs []uint) ([]uint, error) {
	var IDs []productID
	tx := repo.db.Model(&entities.Product{}).Where("id IN ?", productIDs).Find(&IDs)

	ids := []uint{}

	for _, id := range IDs {
		ids = append(ids, id.ID)
	}
	return ids, tx.Error
}
