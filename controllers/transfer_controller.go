package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type TransferController interface {
	TransferGift(ctx *gin.Context)
	GetTransferGift(ctx *gin.Context)
}

type transferController struct {
	giftService     services.GiftService
	accountService  services.AccountService
	transferService services.TransferService
}

func NewTransferController(giftService services.GiftService,
	accountService services.AccountService,
	transferService services.TransferService) TransferController {
	return &transferController{
		giftService:     giftService,
		accountService:  accountService,
		transferService: transferService,
	}
}

func (c *transferController) TransferGift(ctx *gin.Context) {
	transferDTO := ctx.MustGet(configs.TRANSFER_DTO_KEY).(dto.TransferGiftDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	_, err := c.accountService.FindByID(providerID, transferDTO.AccountID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	_, err = c.giftService.FindByID(transferDTO.GiftID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	if err = c.transferService.CreateTransfer(transferDTO, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Transfer gift successfully",
	})
}

func (c *transferController) GetTransferGift(ctx *gin.Context) {
	accountID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	transfers, err := c.transferService.GetTransferByAccount(accountID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"transfers": transfers,
		},
		"error":   nil,
		"message": "Transfer found successfully",
	})
}
