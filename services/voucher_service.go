package services

import (
	"errors"
	"math"

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
	voucher := voucherDTO.ToEntity(providerID)
	return s.voucherRepo.Create(voucher)
}

func (s *voucherService) Update(voucherDTO dto.VoucherDTO, providerID uint, voucherID uint) error {
	voucher := voucherDTO.ToEntity(providerID)
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
	var pagingMetadata viewmodel.PagingMetadata

	count, err := s.voucherRepo.Count(providerID)
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

	vouchers, err := s.voucherRepo.FindAllWithPage(providerID, page, perPage)
	return pagingMetadata, vouchers, err
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
