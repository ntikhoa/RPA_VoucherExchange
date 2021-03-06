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
	FindByIDByAccount(providerID uint, accountID uint, receiptID uint) (entities.Receipt, error)
	Censor(providerID uint, receiptID uint, isApproved bool) error
	FindBetweenDates(providerID uint, fromDate time.Time, toDate time.Time) ([]viewmodel.ReceiptListRes, error)
	FindBetweenDatesWithAccount(providerID uint, accountID uint, fromDate time.Time, toDate time.Time) ([]viewmodel.ReceiptListRes, error)
	FindAllByAccount(
		accountID uint,
		providerID uint,
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

	pagingMetadata, err := paging(s.repo.Count, providerID, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	receipts, err := s.repo.FindAllWithPage(providerID, page, perPage)
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

	if receipt.Account.ProviderID != providerID {
		return receipt, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
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

func (s *receiptService) FindBetweenDates(providerID uint, fromDate time.Time, toDate time.Time) ([]viewmodel.ReceiptListRes, error) {

	if toDate.Before(fromDate) {
		return nil, custom_error.NewBadRequestError("invalid date range")
	}

	receipts, err := s.repo.FindBetweenDates(providerID, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	res := viewmodel.NewSliceReceiptListRes(receipts)

	return res, nil
}

func (s *receiptService) FindAllByAccount(
	accountID uint,
	providerID uint,
	page int,
	perPage int) (viewmodel.PagingMetadata, []viewmodel.ReceiptListRes, error) {

	count, err := s.repo.CountByAccount(providerID, accountID)
	if err != nil {
		return viewmodel.PagingMetadata{}, nil, err
	}

	pagingMetadata, err := paging2(count, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	receipts, err := s.repo.FindAllWithAccountWithPage(providerID, accountID, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	receiptsRes := viewmodel.NewSliceReceiptListRes(receipts)

	return pagingMetadata, receiptsRes, nil
}
func (s *receiptService) FindByIDByAccount(providerID uint, accountID uint, receiptID uint) (entities.Receipt, error) {
	receipt, err := s.repo.FindByID(providerID, receiptID)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return receipt, custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return receipt, err
	}

	if receipt.Account.ProviderID != providerID || receipt.AccountID != accountID {
		return receipt, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return receipt, nil
}

func (s *receiptService) FindBetweenDatesWithAccount(providerID uint, accountID uint, fromDate time.Time, toDate time.Time) ([]viewmodel.ReceiptListRes, error) {
	if toDate.Before(fromDate) {
		return nil, custom_error.NewBadRequestError("invalid date range")
	}

	receipts, err := s.repo.FindBetweenDatesWithAccount(providerID, accountID, fromDate, toDate)
	if err != nil {
		return nil, err
	}

	res := viewmodel.NewSliceReceiptListRes(receipts)

	return res, nil
}
