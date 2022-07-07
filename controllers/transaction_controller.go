package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	FindAll(ctx *gin.Context)
}

type transactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) TransactionController {
	return &transactionController{service: service}
}

func (c *transactionController) FindAll(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	receipts, err := c.service.FindAll(providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"receipts": receipts,
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}
