package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type AccountService interface {
	FindAllWithPage(providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.AccountResponse, error)
	Search(query string, providerID uint) ([]viewmodel.AccountResponse, error)
	FindByID(providerID uint, accountID uint) (viewmodel.AccountResponse, error)
	GetAccountProfile(accountID uint) (entities.Account, error)
}

type accountService struct {
	repo repositories.AccountRepo
}

func NewAccountService(accountRepo repositories.AccountRepo) AccountService {
	return &accountService{
		repo: accountRepo,
	}
}

func (s *accountService) FindAllWithPage(providerID uint,
	page int,
	perPage int) (viewmodel.PagingMetadata, []viewmodel.AccountResponse, error) {
	pagingMetadata := viewmodel.PagingMetadata{}

	pagingMetadata, err := paging(s.repo.Count, providerID, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	accounts, err := s.repo.FindAllWithPage(providerID, page, perPage)
	return pagingMetadata, accounts, err
}

func (s *accountService) Search(query string, providerID uint) ([]viewmodel.AccountResponse, error) {
	accounts, err := s.repo.Search(query, providerID)
	if err != nil {
		return accounts, err
	}
	return accounts, err
}

func (s *accountService) FindByID(providerID uint, accountID uint) (viewmodel.AccountResponse, error) {
	account, err := s.repo.FindByID(accountID, providerID)

	var accountRes viewmodel.AccountResponse

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return accountRes, custom_error.NewNotFoundError("account " + constants.NOT_FOUND_ERROR)
		}
		return accountRes, err
	}

	accountRes = viewmodel.NewAccountResponse(account)

	return accountRes, nil
}

func (s *accountService) GetAccountProfile(accountID uint) (entities.Account, error) {
	account, err := s.repo.GetAccountProfile(accountID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return account, custom_error.NewNotFoundError("account " + constants.NOT_FOUND_ERROR)
		}
		return account, err
	}

	return account, nil
}
