package dto

type ViewExchangeVoucherDTO struct {
	Products []string
	Prices   []uint
}

type ExchangeVoucherDTO struct {
	ViewExchangeVoucherDTO ViewExchangeVoucherDTO
	VoucherID              uint
	CustomerName           string
	CustomerPhone          string
}
