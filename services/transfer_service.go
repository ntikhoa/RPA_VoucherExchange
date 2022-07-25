package services

import (
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
)

type TransferService interface {
	CreateTransfer(dto dto.TransferGiftDTO, providerID uint) error
	GetTransferByAccount(accountID uint, providerID uint) ([]viewmodel.TransferResponse, error)
}

type transferService struct {
	repo repositories.TransferRepo
}

func NewTransferService(repo repositories.TransferRepo) TransferService {
	return &transferService{repo: repo}
}

func (s *transferService) CreateTransfer(dto dto.TransferGiftDTO, providerID uint) error {
	transfer := entities.NewTransfer(dto)
	return s.repo.CreateTransfer(transfer)
}

func (s *transferService) GetTransferByAccount(accountID uint, providerID uint) ([]viewmodel.TransferResponse, error) {
	transfers, err := s.repo.GetTransfersByAccount(accountID, providerID)
	if err != nil {
		return nil, err
	}
	res := viewmodel.NewSliceTransferResponse(transfers)
	return res, nil
}
