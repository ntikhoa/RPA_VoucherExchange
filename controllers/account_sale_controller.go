package controllers

import (
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"github.com/gin-gonic/gin"
)

type AccountSaleController interface {
	GetAccountProfile(ctx *gin.Context)
}

type accountSaleController struct {
	service services.AccountService
}

func NewAccountSaleController(accountService services.AccountService) AccountSaleController {
	return &accountSaleController{
		service: accountService,
	}
}

func (c *accountSaleController) GetAccountProfile(ctx *gin.Context) {
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)

	account, err := c.service.GetAccountProfile(accountID)
	if err != nil {
		abortCustomError(ctx, err)
		return
	}

	res := viewmodel.NewAccountProfileRes(account)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"profile": res,
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}
