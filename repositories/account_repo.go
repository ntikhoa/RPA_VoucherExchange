package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type AccountRepo interface {
	Create(account entities.Account) error
	Update(account entities.Account) error
	FindByUsername(username string) (entities.Account, error)
	FindAccount(accountID uint, providerID uint) (entities.Account, error)
	Count(providerID uint) (int64, error)
	FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Account, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (repo *accountRepo) Create(account entities.Account) error {
	return repo.db.Create(&account).Error
}

func (repo *accountRepo) Update(account entities.Account) error {
	return repo.db.Save(&account).Error
}

func (repo *accountRepo) FindByUsername(username string) (entities.Account, error) {
	account := entities.Account{}
	err := repo.db.
		Where(&entities.Account{Username: username}).
		First(&account).
		Error

	return account, err
}

func (repo *accountRepo) FindAccount(accountID uint, providerID uint) (entities.Account, error) {
	account := entities.Account{}
	err := repo.db.
		Where(&entities.Account{
			Model: gorm.Model{
				ID: accountID,
			},
			ProviderID: providerID,
		}).
		Find(&account).
		Error

	return account, err
}

func (repo *accountRepo) Count(providerID uint) (int64, error) {
	return 10, nil
}

func (repo *accountRepo) FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Account, error) {
	return []entities.Account{}, nil
}
