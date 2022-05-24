package controllers

import (
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type VoucherController interface {
	Create(ctx *gin.Context)
	FindByID(ctx *gin.Context)
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
	voucherDTO := dto.VoucherDTO{}
	err := ctx.ShouldBind(&voucherDTO)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	productIDs := voucherDTO.GetProductIDs()
	err = c.productService.CheckExistence(productIDs)
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}

	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	err = c.voucherService.Create(voucherDTO, providerID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Voucher created successfully.",
	})
}

func (c *voucherController) FindByID(ctx *gin.Context) {
	voucherID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	voucher, err := c.voucherService.FindByID(voucherID, providerID)
	if err != nil {
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    voucher,
		"error":   nil,
		"message": "Voucher found successfully.",
	})
}
