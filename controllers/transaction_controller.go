package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
}

type transactionController struct {
	service services.ReceiptService
}

func NewTransactionController(service services.ReceiptService) TransactionController {
	return &transactionController{service: service}
}

func (c *transactionController) FindAll(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	page := ctx.MustGet(configs.PAGE_QUERY_KEY).(int)
	perPage := ctx.MustGet(configs.PER_PAGE_QUERY_KEY).(int)

	metadata, receipts, err := c.service.FindAll(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"metadata": metadata,
			"receipts": receipts,
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}

func (c *transactionController) FindByID(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	receiptId := ctx.MustGet(configs.ID_PARAM_KEY).(uint)

	receipt, err := c.service.FindByID(providerID, receiptId)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"receipts": receipt,
			"account":  viewmodel.NewAccountResponse(receipt.Account),
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}
