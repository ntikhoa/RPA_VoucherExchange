package services

import (
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type AccountService interface {
	FindAllWithPage(providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.AccountResponse, error)
	Search(query string, providerID uint) ([]viewmodel.AccountResponse, error)
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
