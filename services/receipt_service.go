package services

import (
	"errors"
	"time"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type ReceiptService interface {
	Create(dto dto.ExchangeVoucherDTO, filesName []string, accountID uint) error
	FindAll(providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.ReceiptListRes, error)
	FindByID(providerID uint, receiptID uint) (entities.Receipt, error)
	Censor(providerID uint, receiptID uint, isApproved bool) error
	FindBetweenDates(providerID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error)
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

	pagingMetadata, err := paging(s.repo.Count, providerID, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	receipts, err := s.repo.FindAllWithPage(providerID)
	if err != nil {
		return pagingMetadata, nil, err
	}

	receiptsRes := viewmodel.NewSliceReceiptListRes(receipts)

	return pagingMetadata, receiptsRes, nil
}

func (s *receiptService) FindByID(providerID uint, receiptID uint) (entities.Receipt, error) {
	receipt, err := s.repo.FindByID(providerID, receiptID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return receipt, custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return receipt, err
	}

	return receipt, nil
}

func (s *receiptService) Censor(providerID uint, receiptID uint, isApproved bool) error {
	receipt, err := s.repo.FindByIDWithoutJoin(providerID, receiptID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}

	if receipt.Account.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	statusID := constants.STATUS_REJECTED
	if isApproved {
		statusID = constants.STATUS_APPROVED
	}

	return s.repo.UpdateCensorStatus(receiptID, statusID)
}

func (s *receiptService) FindBetweenDates(providerID uint, fromDate time.Time, toDate time.Time) ([]entities.Receipt, error) {

	if toDate.Before(fromDate) {
		return nil, custom_error.NewBadRequestError("invalid date range")
	}

	receipts, err := s.repo.FindBetweenDates(providerID, fromDate, toDate)
	if err != nil {
		return nil, err
	}
	return receipts, nil
}
