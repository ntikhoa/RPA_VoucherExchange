package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"github.com/gin-gonic/gin"
)

type TransactionSaleController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindBetweenDates(ctx *gin.Context)
}

type transactionSaleController struct {
	service services.ReceiptService
}

func NewTransactionSaleController(service services.ReceiptService) TransactionSaleController {
	return &transactionSaleController{service: service}
}

func (c *transactionSaleController) FindAll(ctx *gin.Context) {
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	page := ctx.MustGet(configs.PAGE_QUERY_KEY).(int)
	perPage := ctx.MustGet(configs.PER_PAGE_QUERY_KEY).(int)

	metadata, receipts, err := c.service.FindAllByAccount(accountID, providerID, page, perPage)
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

func (c *transactionSaleController) FindByID(ctx *gin.Context) {
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	receiptId := ctx.MustGet(configs.ID_PARAM_KEY).(uint)

	receipt, err := c.service.FindByIDByAccount(providerID, accountID, receiptId)
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

func (c *transactionSaleController) FindBetweenDates(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	fromDate := ctx.MustGet(configs.FROM_DATE).(time.Time)
	toDate := ctx.MustGet(configs.TO_DATE).(time.Time)

	receipts, err := c.service.FindBetweenDatesWithAccount(providerID, accountID, fromDate, toDate)
	if err != nil {
		log.Print(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"receipts": receipts,
		},
		"error":   nil,
		"message": "Query completed successfully.",
	})
}
