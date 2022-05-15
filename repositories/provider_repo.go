package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type ProviderRepo interface {
	FindProviderByID(providerID uint) (entities.Provider, error)
}

type providerRepo struct {
	db *gorm.DB
}

func NewProviderRepo(db *gorm.DB) ProviderRepo {
	return &providerRepo{
		db: db,
	}
}

func (repo *providerRepo) FindProviderByID(providerID uint) (entities.Provider, error) {
	provider := entities.Provider{
		Model: gorm.Model{
			ID: providerID,
		},
	}
	err := repo.db.First(&provider).Error
	return provider, err
}
