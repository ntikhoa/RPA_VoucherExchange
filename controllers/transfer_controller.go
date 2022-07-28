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
	TransferGifts(ctx *gin.Context)
	GetTransferGift(ctx *gin.Context)
	AcceptGifts(ctx *gin.Context)
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

func (c *transferController) TransferGifts(ctx *gin.Context) {
	transferGiftsDTO := ctx.MustGet(configs.TRANSFER_DTO_KEY).(dto.CreateTransferGiftsDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	_, err := c.accountService.FindByID(providerID, transferGiftsDTO.AccountID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	var giftIDs []uint
	for _, transfersDTO := range transferGiftsDTO.TransferGiftDTO {
		giftIDs = append(giftIDs, transfersDTO.GiftID)
	}

	err = c.giftService.CheckExistence(providerID, giftIDs)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	if err = c.transferService.CreateTransfers(transferGiftsDTO, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Transfer gifts successfully",
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

func (c *transferController) AcceptGifts(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	accountID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)

	_, err := c.accountService.FindByID(providerID, accountID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	err = c.transferService.AcceptTransfers(accountID)
	if err != nil {
		log.Println()
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Accept gifts successfully",
	})
}
