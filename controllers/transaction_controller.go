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

type TransactionController interface {
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Censor(ctx *gin.Context)
	FindBetweenDates(ctx *gin.Context)
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
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	roleID := ctx.MustGet(configs.ACCOUNT_ROLE_ID_KEY).(uint)

	metadata, receipts, err := c.service.FindAll(providerID, page, perPage, accountID, roleID)
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
	receiptID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	roleID := ctx.MustGet(configs.ACCOUNT_ROLE_ID_KEY).(uint)

	receipt, err := c.service.FindByID(providerID, receiptID, accountID, roleID)
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

func (c *transactionController) Censor(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	receiptID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	isApproved := ctx.MustGet(configs.IS_APPROVED_KEY).(bool)

	if err := c.service.Censor(providerID, receiptID, isApproved); err != nil {
		log.Print(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Transaction censored successfully.",
	})
}

func (c *transactionController) FindBetweenDates(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	roleID := ctx.MustGet(configs.ACCOUNT_ROLE_ID_KEY).(uint)
	fromDate := ctx.MustGet(configs.FROM_DATE).(time.Time)
	toDate := ctx.MustGet(configs.TO_DATE).(time.Time)

	receipts, err := c.service.FindBetweenDates(providerID, accountID, roleID, fromDate, toDate)
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
