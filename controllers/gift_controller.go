package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type GiftController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteGifts(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Search(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type giftController struct {
	giftService services.GiftService
}

func NewGiftController(giftService services.GiftService) GiftController {
	return &giftController{
		giftService: giftService,
	}
}

func (c *giftController) Create(ctx *gin.Context) {
	giftDTO := ctx.MustGet(configs.GIFT_DTO_KEY).(dto.GiftDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.giftService.Create(giftDTO, providerID); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"data":    nil,
		"error":   nil,
		"message": "Gift created successfully.",
	})
}

func (c *giftController) Update(ctx *gin.Context) {

	giftDTO := ctx.MustGet(configs.GIFT_DTO_KEY).(dto.GiftDTO) //from gift validation middlewares
	giftID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.giftService.Update(giftDTO, providerID, giftID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Gift updated successfully.",
	})
}

func (c *giftController) Delete(ctx *gin.Context) {

	giftID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	if err := c.giftService.DeleteByID(giftID, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Gift deleted successfully.",
	})
}

func (c *giftController) FindAll(ctx *gin.Context) {
	page := ctx.MustGet(configs.PAGE_QUERY_KEY).(int)
	perPage := ctx.MustGet(configs.PER_PAGE_QUERY_KEY).(int)

	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	metadata, gifts, err := c.giftService.FindAllWithPage(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"metadata": metadata,
			"gifts":    gifts,
		},
		"error":   nil,
		"message": "Gift found successfully.",
	})
}

func (c *giftController) FindByID(ctx *gin.Context) {

	giftID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	gift, err := c.giftService.FindByID(giftID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"gift": gift,
		},
		"error":   nil,
		"message": "Gift found successfully.",
	})
}

func (c *giftController) GetAll(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	gifts, err := c.giftService.GetAll(providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"gifts": gifts,
		},
		"error":   nil,
		"message": "Gift created successfully.",
	})
}

func (c *giftController) DeleteGifts(ctx *gin.Context) {
	giftIDs := ctx.MustGet(configs.PAYLOAD_DTO_KEY).([]uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.giftService.DeleteByIDs(giftIDs, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Gifts deleted successfully.",
	})
}

func (c *giftController) Search(ctx *gin.Context) {
	query := ctx.MustGet(configs.SEARCH_QUERY_KEY).(string)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	gifts, err := c.giftService.Search(query, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"gifts": gifts,
		},
		"error":   nil,
		"message": "Gifts found successfully"})
}
