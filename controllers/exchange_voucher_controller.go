package controllers

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
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
	products := ctx.MustGet(configs.EDITED_PRODUCTS_KEY).([]string)
	prices := ctx.MustGet(configs.EDITED_PRICES_KEY).([]uint)

	vouchers, err := c.service.ExchangeVoucher(providerID, products, prices)
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
	products := ctx.MustGet(configs.EDITED_PRODUCTS_KEY).([]string)
	prices := ctx.MustGet(configs.EDITED_PRICES_KEY).([]uint)

	vouchers, err := c.service.ExchangeVoucher(providerID, products, prices)
	if err != nil {
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

	ctx.JSON(http.StatusOK, gin.H{
		"data": vouchers,
	})
}

// func debugPrintParseStrReq(value []string) {
// 	for _, value := range value {
// 		log.Println(value)
// 	}
// }

// func debugPrintParseUintReq(value []uint) {
// 	for _, value := range value {
// 		log.Println(value)
// 	}
// }
