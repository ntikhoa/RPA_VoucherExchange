package controllers

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type ExchangeVoucherController interface {
	ViewExchangeVoucher(ctx *gin.Context)
	ExchangeVoucher(ctx *gin.Context)
}

type exchangeVoucherController struct {
	service services.ExchangeVoucherService
}

func NewExchangeVoucherController(service services.ExchangeVoucherService) ExchangeVoucherController {
	return &exchangeVoucherController{
		service: service,
	}
}

func (c *exchangeVoucherController) ViewExchangeVoucher(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	viewVoucherExchangeDTO := ctx.MustGet(configs.VIEW_EXCHANGE_VOUCHER_DTO_KEY).(dto.ViewExchangeVoucherDTO)

	vouchers, err := c.service.ViewExchangeVoucher(providerID, viewVoucherExchangeDTO)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": vouchers,
	})
}

func (c *exchangeVoucherController) ExchangeVoucher(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	files := ctx.MustGet(configs.RECEIPT_IMAGE_FILES_KEY).([]*multipart.FileHeader)
	dto := ctx.MustGet(configs.EXCHANGE_VOUCHER_DTO).(dto.ExchangeVoucherDTO)

	if err := c.service.ExchangeVoucher(providerID, dto.ViewExchangeVoucherDTO, dto.VoucherID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	var filesName []string
	s3Service := services.NewImageService()
	for _, file := range files {
		fileName, err := s3Service.UploadObject(file)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		filesName = append(filesName, fileName)
	}
}
