package services

import (
	"math"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type AccountService interface {
	FindAllWithPage(providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.AccountResponse, error)
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

	count, err := s.repo.Count(providerID)
	if err != nil {
		return pagingMetadata, nil, err
	}
	d := float64(count) / float64(perPage)
	totalPages := int(math.Ceil(d))
	if page > totalPages {
		return pagingMetadata, nil, custom_error.NewNotFoundError(constants.EXHAUSTED_ERROR)
	}

	pagingMetadata = viewmodel.PagingMetadata{
		Page:         page,
		PerPage:      perPage,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}
	accounts, err := s.repo.FindAllWithPage(providerID, page, perPage)
	return pagingMetadata, accounts, err
}
