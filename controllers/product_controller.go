package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	FindAllProduct(ctx *gin.Context)
	FindProductByID(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) CreateProduct(ctx *gin.Context) {
	productDTO := ctx.MustGet(configs.PRODUCT_DTO_KEY).(dto.ProductDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.productService.CreateProduct(productDTO, providerID); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Product created successfully.",
	})
}

func (c *productController) UpdateProduct(ctx *gin.Context) {

	productDTO := ctx.MustGet(configs.PRODUCT_DTO_KEY).(dto.ProductDTO) //from product validation middlewares
	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.productService.UpdateProduct(productDTO, providerID, productID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Product updated successfully.",
	})
}

func (c *productController) DeleteProduct(ctx *gin.Context) {

	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	if err := c.productService.DeleteProductByID(productID, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "Product deleted successfully.",
	})
}

func (c *productController) FindAllProduct(ctx *gin.Context) {
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
	metadata, products, err := c.productService.FindAllProductWithPage(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"metadata": metadata,
			"products": products,
		},
		"error":   nil,
		"message": "Product found successfully.",
	})
}

func (c *productController) FindProductByID(ctx *gin.Context) {

	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	product, err := c.productService.FindProductByID(productID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data": gin.H{
			"product": product,
		},
		"error":   nil,
		"message": "Product found successfully.",
	})
}
