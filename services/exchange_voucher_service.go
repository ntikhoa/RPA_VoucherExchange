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

type ExchangeVoucherService interface {
	ViewExchangeVoucher(providerID uint, viewVoucherExchangeDTO dto.ViewExchangeVoucherDTO) ([]entities.Voucher, error)
	ExchangeVoucher(providerID uint, viewVoucherExchangeDTO dto.ViewExchangeVoucherDTO, voucherID uint) error
	checkVoucherRule(dto dto.ViewExchangeVoucherDTO,
		vouchers []entities.Voucher,
	) []entities.Voucher
}

type exchangeVoucherService struct {
	voucherRepo repositories.VoucherRepo
}

func NewExchangeVoucherService(voucherRepo repositories.VoucherRepo) ExchangeVoucherService {
	return &exchangeVoucherService{
		voucherRepo: voucherRepo,
	}
}

func (s *exchangeVoucherService) ViewExchangeVoucher(
	providerID uint,
	dto dto.ViewExchangeVoucherDTO,
) ([]entities.Voucher, error) {
	vouchers, err := s.voucherRepo.FindVoucherExchange(providerID, dto.Products)
	if err != nil {
		return nil, err
	}

	resVoucher := s.checkVoucherRule(dto, vouchers)
	return resVoucher, nil
}

func (s *exchangeVoucherService) ExchangeVoucher(providerID uint,
	viewVoucherExchangeDTO dto.ViewExchangeVoucherDTO,
	voucherID uint,
) error {
	voucher, err := s.voucherRepo.FindByID(voucherID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}
	if voucher.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	resVoucher := s.checkVoucherRule(viewVoucherExchangeDTO, []entities.Voucher{voucher})
	if len(resVoucher) != 1 {
		return custom_error.NewConflictError(constants.INVALID_VOUCHER_ERROR)
	}
	return nil
}

func (s *exchangeVoucherService) checkVoucherRule(
	dto dto.ViewExchangeVoucherDTO,
	vouchers []entities.Voucher,
) []entities.Voucher {

	var resVoucher []entities.Voucher
	productPrice := make(map[string]uint)
	for i, product := range dto.Products {
		productPrice[product] = dto.Prices[i]
	}
	for _, voucher := range vouchers {
		totalPrices := uint(0)
		for _, product := range voucher.Products {
			price, ok := productPrice[product.ProductName]
			if ok {
				totalPrices += price
			}
		}
		if totalPrices >= voucher.TotalPriceMin {
			if voucher.TotalPriceMax != 0 {
				if totalPrices <= voucher.TotalPriceMax {
					resVoucher = append(resVoucher, voucher)
				}
			} else {
				resVoucher = append(resVoucher, voucher)
			}
		}
	}
	return resVoucher
}
