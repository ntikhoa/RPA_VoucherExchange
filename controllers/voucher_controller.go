package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type VoucherController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Publish(ctx *gin.Context)
}

type voucherController struct {
	voucherService services.VoucherService
	productService services.ProductService
}

func NewVoucherController(voucherService services.VoucherService,
	productService services.ProductService) VoucherController {
	return &voucherController{
		voucherService: voucherService,
		productService: productService,
	}
}

func (c *voucherController) Create(ctx *gin.Context) {
	voucherDTO := ctx.MustGet(configs.VOUCHER_DTO_KEY).(dto.VoucherDTO)

	productIDs := voucherDTO.ProductIDs
	if err := c.productService.CheckExistence(productIDs); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	if err := c.voucherService.Create(voucherDTO, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"data":    nil,
		"error":   nil,
		"message": "Voucher created successfully.",
	})
}

func (c *voucherController) Update(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	voucherID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	voucherDTO := ctx.MustGet(configs.VOUCHER_DTO_KEY).(dto.VoucherDTO)

	productIDs := voucherDTO.ProductIDs
	if err := c.productService.CheckExistence(productIDs); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	if err := c.voucherService.Update(voucherDTO, providerID, voucherID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Voucher updated successfully.",
	})
}

func (c *voucherController) FindByID(ctx *gin.Context) {
	voucherID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	voucher, err := c.voucherService.FindByID(voucherID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    voucher,
		"error":   nil,
		"message": "Voucher found successfully.",
	})
}

func (c *voucherController) FindAll(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	page := ctx.MustGet(configs.PAGE_QUERY_KEY).(int)
	perPage := 2

	metadata, vouchers, err := c.voucherService.FindAllWithPage(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"metadata": metadata,
			"vouchers": vouchers,
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}

func (c *voucherController) Delete(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	voucherID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)

	if err := c.voucherService.Delete(providerID, voucherID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Voucher deleted successfully.",
	})
}

func (c *voucherController) Publish(ctx *gin.Context) {
	publishedDTO := ctx.MustGet(configs.PUBLISHED_DTO_KEY).(dto.PublishedDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	err := c.voucherService.Publish(providerID, publishedDTO)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Voucher published successfully.",
	})
}
