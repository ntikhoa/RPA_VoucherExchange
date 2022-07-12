package controllers

import (
	"errors"
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

func (c *transactionController) Censor(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	receiptId := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	isApproved := ctx.MustGet(configs.IS_APPROVED_KEY).(bool)

	if err := c.service.Censor(providerID, receiptId, isApproved); err != nil {
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
	//ddmmyyyy
	layout := "02012006"
	fromDate_string := ctx.MustGet(configs.FROM_DATE).(string)
	toDate_string := ctx.MustGet(configs.TO_DATE).(string)

	fromDate, err := time.Parse(layout, fromDate_string)
	if err != nil {
		log.Print(err)
		abortCustomError(ctx, err)
		return
	}
	toDate, err := time.Parse(layout, toDate_string)
	if err != nil {
		log.Print(err)
		abortCustomError(ctx, err)
		return
	}
	// fromDate.Format("2006-01-02 15:04:05")
	// toDate.Format("2006-01-02 15:04:05")

	// log.Println("Transaction from:", fromDate)
	// log.Println("Transaction till:", toDate)

	if toDate.Before(fromDate) {
		err = errors.New("To date < From date.")
		log.Print(err)
		abortCustomError(ctx, err)
		return
	}
	receipts, err := c.service.FindBetweenDates(providerID, fromDate, toDate)
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
