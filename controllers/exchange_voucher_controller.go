package controllers

import (
	"log"
	"mime/multipart"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type ExchangeVoucherController interface {
	ExchangeVoucher(ctx *gin.Context)
}

type exchangeVoucherController struct {
}

func NewExchangeVoucherController() ExchangeVoucherController {
	return &exchangeVoucherController{}
}

func (c *exchangeVoucherController) ExchangeVoucher(ctx *gin.Context) {
	files := ctx.MustGet(configs.RECEIPT_IMAGE_FILES_KEY).([]*multipart.FileHeader)
	ocrProducts := ctx.MustGet(configs.ORIGINAL_PRODUCTS_KEY).([]string)
	ocrPrices := ctx.MustGet(configs.ORIGINAL_PRICES_KEY).([]uint)
	products := ctx.MustGet(configs.EDITED_PRODUCTS_KEY).([]string)
	prices := ctx.MustGet(configs.EDITED_PRICES_KEY).([]uint)

	s3Service := services.NewImageService()
	for _, file := range files {
		fileName, err := s3Service.UploadObject(file)
		if err != nil {
			panic(err)
		}
		log.Println(fileName)
	}

	debugPrintParseStrReq(ocrProducts)
	debugPrintParseUintReq(ocrPrices)
	debugPrintParseStrReq(products)
	debugPrintParseUintReq(prices)
}

func debugPrintParseStrReq(value []string) {
	for _, value := range value {
		log.Println(value)
	}
}

func debugPrintParseUintReq(value []uint) {
	for _, value := range value {
		log.Println(value)
	}
}
