package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type AccountController interface {
	GetAccount(ctx *gin.Context)
}

type accountController struct {
	service services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return &accountController{
		service: accountService,
	}
}

func (c *accountController) GetAccount(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	page := ctx.MustGet(configs.PAGE_QUERY_KEY).(int)
	perPage := ctx.MustGet(configs.PER_PAGE_QUERY_KEY).(int)

	metadata, accounts, err := c.service.FindAllWithPage(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"metadata": metadata,
			"accounts": accounts,
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}
