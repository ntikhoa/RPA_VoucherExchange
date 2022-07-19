package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type AccountRepo interface {
	Create(account entities.Account) error
	Update(account entities.Account) error
	FindByUsername(username string) (entities.Account, error)
	FindAccount(accountID uint, providerID uint) (entities.Account, error)
	Count(providerID uint) (int64, error)
	FindAllWithPage(providerID uint, page int, perPage int) ([]viewmodel.AccountResponse, error)
	Search(query string, providerID uint) ([]viewmodel.AccountResponse, error)
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
	return repo.db.Omit("created_at").Save(&account).Error
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
	var count int64
	err := repo.db.
		Model(&entities.Account{ProviderID: providerID}).
		Count(&count).
		Error
	return count, err
}

func (repo *accountRepo) FindAllWithPage(providerID uint,
	page int,
	perPage int) ([]viewmodel.AccountResponse, error) {

	var accountsRes []viewmodel.AccountResponse
	err := repo.db.
		Model(&entities.Account{}).
		Preload("Role").
		Where(&entities.Account{ProviderID: providerID}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&accountsRes).
		Error

	return accountsRes, err
}

func (repo *accountRepo) Search(query string, providerID uint) ([]viewmodel.AccountResponse, error) {
	var accountsRes []viewmodel.AccountResponse

	query = "%" + query + "%"
	err := repo.db.
		Model(&entities.Account{}).
		Preload("Role").
		Where(&entities.Account{ProviderID: providerID}).
		Where("username LIKE ? OR name LIKE ?", query, query).
		Distinct().
		Find(&accountsRes).
		Error
	return accountsRes, err
}
