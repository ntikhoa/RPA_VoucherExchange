package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type GiftController interface {
	CreateGift(ctx *gin.Context)
	UpdateGift(ctx *gin.Context)
	DeleteGift(ctx *gin.Context)
	FindAllGift(ctx *gin.Context)
	FindGiftByID(ctx *gin.Context)
}

type giftController struct {
	giftService services.GiftService
}

func NewGiftController(giftService services.GiftService) GiftController {
	return &giftController{
		giftService: giftService,
	}
}

func (c *giftController) CreateGift(ctx *gin.Context) {
	gift := ctx.MustGet(configs.GIFT_KEY).(entities.Gift)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	gift.ProviderID = providerID
	if err := c.giftService.CreateGift(gift); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Gift created successfully.",
	})
}

func (c *giftController) UpdateGift(ctx *gin.Context) {
	giftID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	gift := ctx.MustGet(configs.GIFT_KEY).(entities.Gift)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	gift.Model.ID = giftID
	gift.ProviderID = providerID
	if err := c.giftService.UpdateGift(gift); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Gift updated successfully.",
	})
}

func (c *giftController) DeleteGift(ctx *gin.Context) {
	giftID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	if err := c.giftService.DeleteGiftByID(giftID, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Gift deleted successfully.",
	})
}

func (c *giftController) FindAllGift(ctx *gin.Context) {
	value := ctx.Request.URL.Query()
	pageQuery := value["page"]
	if len(pageQuery) != 1 {
		err := errors.New("page query is required")
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pageConv, err := strconv.ParseInt(pageQuery[0], 10, 64)
	if err != nil || pageConv <= 0 {
		err := errors.New("invalid page query, page should be greater than or equal to 1")
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	page := int(pageConv)
	perPage := 2

	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	metadata, gifts, err := c.giftService.FindAllGiftWithPage(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"metadata": metadata,
			"gifts":    gifts,
		},
		"error":   nil,
		"message": "Gift found successfully.",
	})
}

func (c *giftController) FindGiftByID(ctx *gin.Context) {

	giftID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	gift, err := c.giftService.FindGiftByID(giftID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"gift": gift,
		},
		"error":   nil,
		"message": "Gift found successfully.",
	})
}
