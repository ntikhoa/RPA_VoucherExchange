package services

import (
	"math"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type ReceiptService interface {
	Create(dto dto.ExchangeVoucherDTO, filesName []string, accountID uint) error
	FindAll(providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.ReceiptListRes, error)
}

type receiptService struct {
	repo repositories.ReceiptRepo
}

func NewReceiptService(repo repositories.ReceiptRepo) ReceiptService {
	return &receiptService{
		repo: repo,
	}
}

func (s *receiptService) Create(dto dto.ExchangeVoucherDTO, filesName []string, accountID uint) error {
	receipt := entities.NewReceipt(dto, filesName, accountID)
	return s.repo.Create(receipt)
}

func (s *receiptService) FindAll(providerID uint,
	page int,
	perPage int) (viewmodel.PagingMetadata, []viewmodel.ReceiptListRes, error) {

	var pagingMetadata viewmodel.PagingMetadata

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

	receipts, err := s.repo.FindAllWithPage(providerID)
	if err != nil {
		return pagingMetadata, nil, err
	}

	receiptsRes := viewmodel.NewSliceReceiptListRes(receipts)

	return pagingMetadata, receiptsRes, nil
}
