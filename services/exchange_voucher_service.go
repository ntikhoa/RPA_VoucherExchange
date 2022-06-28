package services

import (
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
)

type ExchangeVoucherService interface {
	ExchangeVoucher(providerID uint, products []string, prices []uint) ([]entities.Voucher, error)
}

type exchangeVoucherService struct {
	voucherRepo repositories.VoucherRepo
}

func NewExchangeVoucherService(voucherRepo repositories.VoucherRepo) ExchangeVoucherService {
	return &exchangeVoucherService{
		voucherRepo: voucherRepo,
	}
}

func (s *exchangeVoucherService) ExchangeVoucher(providerID uint, products []string, prices []uint) ([]entities.Voucher, error) {
	var resVoucher []entities.Voucher

	vouchers, err := s.voucherRepo.FindVoucherExchange(providerID, products)
	if err != nil {
		return vouchers, err
	}
	productPrice := make(map[string]uint)
	for i, product := range products {
		productPrice[product] = prices[i]
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
	return resVoucher, nil
}
