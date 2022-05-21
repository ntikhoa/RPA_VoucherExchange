package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/dto"
)

type VoucherService interface {
	CreateVoucher(voucherDTO dto.VoucherDTO) error
}

type voucherService struct {
}

func NewVoucherService() VoucherService {
	return &voucherService{}
}

func (s *voucherService) CreateVoucher(voucherDTO dto.VoucherDTO) error {
	// s.productRepo.FindProductByID()

	return errors.New("dummy")
}
