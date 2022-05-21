package controllers

import (
	"net/http"

	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type VoucherController interface {
	CreateVoucher(ctx *gin.Context)
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

func (c *voucherController) CreateVoucher(ctx *gin.Context) {
	voucherDTO := dto.VoucherDTO{}
	err := ctx.ShouldBind(&voucherDTO)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	productIDs := voucherDTO.GetVoucherProducts()
	err = c.productService.CheckProductsExist(productIDs)
	if err != nil {
		ctx.AbortWithError(http.StatusConflict, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": voucherDTO,
	})
}
