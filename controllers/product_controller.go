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
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindByID(ctx *gin.Context)

	fetchProductIDParam(ctx *gin.Context) (uint, error)
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

func (c *productController) Create(ctx *gin.Context) {
	product := ctx.MustGet(configs.PRODUCT_KEY).(entities.Product)
	product.ProviderID = globalProviderID
	if err := c.productService.Create(product); err != nil {
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

func (c *productController) Update(ctx *gin.Context) {
	productID, err := c.fetchProductIDParam(ctx)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product := ctx.MustGet(configs.PRODUCT_KEY).(entities.Product)
	product.Model.ID = productID
	product.ProviderID = globalProviderID
	if err := c.productService.Update(product); err != nil {
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

func (c *productController) Delete(ctx *gin.Context) {
	productID, err := c.fetchProductIDParam(ctx)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := c.productService.DeleteByID(productID, globalProviderID); err != nil {
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

func (c *productController) FindAll(ctx *gin.Context) {
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

	metadata, products, err := c.productService.FindAllWithPage(globalProviderID, page, perPage)
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

func (c *productController) FindByID(ctx *gin.Context) {
	productID, err := c.fetchProductIDParam(ctx)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	product, err := c.productService.FindByID(productID, globalProviderID)
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

//@return ProductID, error
//if fetch successfully, error = nil
//if there is an error, ProductID = 0
func (c *productController) fetchProductIDParam(ctx *gin.Context) (uint, error) {
	productIDParam, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid product ID")
	}
	productID := uint(productIDParam)
	return productID, nil
}
