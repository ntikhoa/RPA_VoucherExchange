package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(product entities.Product) error
	Update(product entities.Product) error
	DeleteById(productID uint) error
	FindAll(providerID uint) ([]entities.Product, error)
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
	return repo.db.Save(&product).Error
}

func (repo *productRepo) DeleteById(productID uint) error {
	return repo.db.Delete(&entities.Product{}, productID).Error
}

func (repo *productRepo) FindAll(providerID uint) ([]entities.Product, error) {
	products := []entities.Product{}
	tx := repo.db.Where("provider_id = ?", providerID).Find(&products)
	return products, tx.Error
}
