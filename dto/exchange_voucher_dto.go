package dto

type ViewExchangeVoucherDTO struct {
	Products []string `json:"products" binding:"required"`
	Prices   []uint   `json:"prices" binding:"required"`
}

type ExchangeVoucherDTO struct {
	ViewExchangeVoucherDTO ViewExchangeVoucherDTO
	TransactionID          string
	VoucherID              uint
	CustomerName           string
	CustomerPhone          string
}
