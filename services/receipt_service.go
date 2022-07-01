package services

import (
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/repositories"
)

type ReceiptService interface {
	Create(dto dto.ExchangeVoucherDTO, filesName []string, accountID uint) error
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
	receipt := dto.ToEntitiy(filesName, accountID)
	return s.repo.Create(receipt)
}
