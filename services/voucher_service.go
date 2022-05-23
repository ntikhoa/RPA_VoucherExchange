package services

import (
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/repositories"
)

type VoucherService interface {
	Create(voucherDTO dto.VoucherDTO, providerID uint) error
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
