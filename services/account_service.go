package services

import (
	"log"

	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type AccountService interface {
	FindAllWithPage(providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.AccountResponse, error)
	FindByUserOrName(query string, providerID uint) ([]entities.Account, error)
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

func (s *accountService) FindByUserOrName(query string, providerID uint) ([]entities.Account, error) {
	accounts, err := s.repo.FindByUserOrName(query, providerID)
	if err != nil {
		log.Print(err)
		return accounts, err
	}
	return accounts, err
}
