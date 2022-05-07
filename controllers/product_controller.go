package controllers

import (
	"errors"
	"net/http"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(ctx *gin.Context) int
	// Update(ctx *gin.Context) error
	// Delete(ctx *gin.Context) error
	// FindAll(ctx *gin.Context) ([]entities.Product, error)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) Create(ctx *gin.Context) int {
	test := &dto.TestDTO{}
	err := ctx.ShouldBind(&test)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return -1
	}

	if test.Secret != "ntikhoa" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("secret does not match"))
		return -1
	}

	providerID, ok := ctx.Get(constants.PROVIDER_ID_KEY)
	if !ok {
		ctx.AbortWithError(http.StatusNotFound, errors.New("does not found key"))
		return -1
	}
	return providerID.(int)
}

// func (c *productController) Update(ctx *gin.Context) error {

// }

// func (c *productController) Delete(ctx *gin.Context) error {

// }

// func (c *productController) FindAll(ctx *gin.Context) ([]entities.Product, error) {

// }
