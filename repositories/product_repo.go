package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ProductRepo interface {
	Create(product entities.Product) error
	Update(product entities.Product) error
	DeleteByID(productID uint) error
	FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Product, error)
	FindByID(productID uint) (entities.Product, error)

	GetCount(providerID uint) (int64, error)
	CheckExistence(productIDs []uint) ([]uint, error)
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

func (repo *productRepo) GetCount(providerID uint) (int64, error) {
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

func (repo *productRepo) CheckExistence(productIDs []uint) ([]uint, error) {
	var IDs []productID
	tx := repo.db.Model(&entities.Product{}).Where("id IN ?", productIDs).Find(&IDs)

	ids := []uint{}

	for _, id := range IDs {
		ids = append(ids, id.ID)
	}
	return ids, tx.Error
}
