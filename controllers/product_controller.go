package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	custom_error "github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

var (
	globalProviderID = uint(1)
)

type ProductController interface {
	CreateProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	FindAllProduct(ctx *gin.Context)
	FindProductByID(ctx *gin.Context)

	abortAndCheckError(ctx *gin.Context, err error)
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
	product := ctx.MustGet(configs.PRODUCT_KEY).(entities.Product)
	product.ProviderID = globalProviderID
	if err := c.productService.CreateProduct(product); err != nil {
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

	productID := ctx.MustGet(configs.ID_PARAM_KEY).(uint)
	product := ctx.MustGet(configs.PRODUCT_KEY).(entities.Product)
	product.Model.ID = productID
	product.ProviderID = globalProviderID
	if err := c.productService.UpdateProduct(product); err != nil {
		log.Println(err)
		c.abortAndCheckError(ctx, err)
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
	if err := c.productService.DeleteProductByID(productID, globalProviderID); err != nil {
		log.Println(err)
		c.abortAndCheckError(ctx, err)
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

	metadata, products, err := c.productService.FindAllProductWithPage(globalProviderID, page, perPage)
	if err != nil {
		log.Println(err)
		c.abortAndCheckError(ctx, err)
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
	product, err := c.productService.FindProductByID(productID, globalProviderID)
	if err != nil {
		log.Println(err)
		c.abortAndCheckError(ctx, err)
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

//check for custom error and abort with the coresponded message
//use for checking custom message sent from the services layer
func (c *productController) abortAndCheckError(ctx *gin.Context, err error) {
	switch err.(type) {
	case *custom_error.AuthorizedError:
		ctx.AbortWithError(http.StatusForbidden, errors.New("you don't have permission to delete this data"))
	case *custom_error.NotFoundError:
		ctx.AbortWithError(http.StatusNotFound, errors.New("data does not exist"))
	case *custom_error.ExhaustedError:
		ctx.AbortWithError(http.StatusNotFound, errors.New("data exhausted, page number exceeds maximum pages"))
	default:
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}
