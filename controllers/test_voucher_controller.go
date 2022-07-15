package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type TestVoucher interface {
	TestVoucher(ctx *gin.Context)
}

type testVoucher struct {
	evService services.ExchangeVoucherService
}

func NewTestVoucher(evService services.ExchangeVoucherService) TestVoucher {
	return &testVoucher{
		evService: evService,
	}
}

func (c *testVoucher) TestVoucher(ctx *gin.Context) {
	testExchangeDTO := ctx.MustGet(configs.TEST_EXCHANGE_VOUCHER_DTO_KEY).(dto.TestExchangeVoucherDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	voucherIDs, err := c.evService.TestExchangeVoucher(providerID, testExchangeDTO)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"satisfied_voucher_id": voucherIDs,
	})
}
