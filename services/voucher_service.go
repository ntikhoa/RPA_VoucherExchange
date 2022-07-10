package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type VoucherService interface {
	Create(voucherDTO dto.VoucherDTO, providerID uint) error
	Update(voucherDTO dto.VoucherDTO, providerID uint, voucherID uint) error
	FindByID(voucherID uint, providerID uint) (entities.Voucher, error)
	FindAllWithPage(
		providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []viewmodel.VoucherResponse, error)
	FindByName(
		voucherName string,
		providerID uint) ([]viewmodel.VoucherResponse, error)
	Delete(providerID uint, voucherID uint) error
	Publish(providerID uint, publishDTO dto.PublishedDTO) error
}

type voucherService struct {
	voucherRepo repositories.VoucherRepo
}

func NewVoucherService(voucherRepo repositories.VoucherRepo) VoucherService {
	return &voucherService{
		voucherRepo: voucherRepo,
	}
}

func (s *voucherService) Create(voucherDTO dto.VoucherDTO, providerID uint) error {
	voucher := entities.NewVoucher(voucherDTO, providerID)
	return s.voucherRepo.Create(voucher)
}

func (s *voucherService) Update(voucherDTO dto.VoucherDTO, providerID uint, voucherID uint) error {
	voucher := entities.NewVoucher(voucherDTO, providerID)
	voucher.Model.ID = voucherID

	//authorize
	fetchedVoucher, err := s.voucherRepo.FindByID(voucherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}
	if fetchedVoucher.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	//Update
	fetchedVoucher.Gift.GiftName = voucherDTO.Gift
	voucher.Gift = fetchedVoucher.Gift
	return s.voucherRepo.Update(voucher)
}

func (s *voucherService) FindByID(voucherID uint, providerID uint) (entities.Voucher, error) {
	voucher, err := s.voucherRepo.FindByID(voucherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return voucher, custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return voucher, err
	}

	if voucher.ProviderID != providerID {
		return voucher, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return voucher, nil
}

func (s *voucherService) FindAllWithPage(
	providerID uint,
	page int,
	perPage int,
) (viewmodel.PagingMetadata, []viewmodel.VoucherResponse, error) {

	pagingMetadata, err := paging(s.voucherRepo.Count, providerID, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	vouchers, err := s.voucherRepo.FindAllWithPage(providerID, page, perPage)
	return pagingMetadata, vouchers, err
}

func (s *voucherService) FindByName(voucherName string, providerID uint) ([]viewmodel.VoucherResponse, error) {
	reponses, err := s.voucherRepo.FindByName(voucherName, providerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return reponses, custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return reponses, err
	}

	return reponses, err
}

func (s *voucherService) Delete(providerID uint, voucherID uint) error {
	voucher, err := s.voucherRepo.FindByID(voucherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	if voucher.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return s.voucherRepo.Delete(voucherID)
}

func (s *voucherService) Publish(providerID uint, publishDTO dto.PublishedDTO) error {
	voucher, err := s.voucherRepo.FindByID(publishDTO.VoucherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}

	if voucher.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return s.voucherRepo.Publish(publishDTO.VoucherID, publishDTO.Published)
}
