package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteProducts(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Search(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) ProductController {
	return &productController{
		productService: productService,
	}
}

func (c *productController) Create(ctx *gin.Context) {
	productDTO := ctx.MustGet(configs.PRODUCT_DTO_KEY).(dto.ProductDTO)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.productService.Create(productDTO, providerID); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"data":    nil,
		"error":   nil,
		"message": "Product created successfully.",
	})
}

func (c *productController) Update(ctx *gin.Context) {

	productDTO := ctx.MustGet(configs.PRODUCT_DTO_KEY).(dto.ProductDTO) //from product validation middlewares
	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.productService.Update(productDTO, providerID, productID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Product updated successfully.",
	})
}

func (c *productController) Delete(ctx *gin.Context) {

	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	if err := c.productService.DeleteByID(productID, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Product deleted successfully.",
	})
}

func (c *productController) FindAll(ctx *gin.Context) {
	page := ctx.MustGet(configs.PAGE_QUERY_KEY).(int)
	perPage := ctx.MustGet(configs.PER_PAGE_QUERY_KEY).(int)

	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	metadata, products, err := c.productService.FindAllWithPage(providerID, page, perPage)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"metadata": metadata,
			"products": products,
		},
		"error":   nil,
		"message": "Product found successfully.",
	})
}

func (c *productController) FindByID(ctx *gin.Context) {

	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	product, err := c.productService.FindByID(productID, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"product": product,
		},
		"error":   nil,
		"message": "Product found successfully.",
	})
}

func (c *productController) GetAll(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	products, err := c.productService.GetAll(providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"products": products,
		},
		"error":   nil,
		"message": "Product created successfully.",
	})
}

func (c *productController) DeleteProducts(ctx *gin.Context) {
	productIDs := ctx.MustGet(configs.PAYLOAD_DTO_KEY).([]uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	if err := c.productService.DeleteByIDs(productIDs, providerID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Products deleted successfully.",
	})
}

func (c *productController) Search(ctx *gin.Context) {
	query := ctx.MustGet(configs.SEARCH_QUERY_KEY).(string)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)

	products, err := c.productService.Search(query, providerID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"products": products,
		},
		"error":   nil,
		"message": "Products found successfully"})
}
