package services

import "github.com/RPA_VoucherExchange/repositories"

type ExchangeVoucherService interface {
	ExchangeVoucher(products []string, prices []string)
}

type exchangeVoucherService struct {
	productRepo repositories.ProductRepo
}

func NewExchangeVoucherService(productRepo repositories.ProductRepo) ExchangeVoucherService {
	return &exchangeVoucherService{
		productRepo: productRepo,
	}
}

func (s *exchangeVoucherService) ExchangeVoucher(products []string, prices []string) {

}
