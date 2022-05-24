package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	"gorm.io/gorm"
)

type VoucherService interface {
	Create(voucherDTO dto.VoucherDTO, providerID uint) error
	FindByID(voucherID uint, providerID uint) (entities.Voucher, error)
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

func (s *voucherService) FindByID(voucherID uint, providerID uint) (entities.Voucher, error) {
	voucher, err := s.voucherRepo.FindByID(voucherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return voucher, &custom_error.NotFoundError{}
		}
		return voucher, err
	}

	if voucher.ProviderID != providerID {
		return voucher, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return s.voucherRepo.FindByID(voucherID)
}
