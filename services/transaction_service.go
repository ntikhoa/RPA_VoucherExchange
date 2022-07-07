package services

import (
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type TransactionService interface {
	FindAll(providerID uint) ([]viewmodel.ReceiptListRes, error)
}

type transactionService struct {
	repo repositories.TransactionRepo
}

func NewTransactionService(repo repositories.TransactionRepo) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) FindAll(providerID uint) ([]viewmodel.ReceiptListRes, error) {
	receipts, err := s.repo.FindAll(providerID)
	if err != nil {
		return nil, err
	}

	receiptsRes := viewmodel.NewSliceReceiptListRes(receipts)

	return receiptsRes, nil
}
