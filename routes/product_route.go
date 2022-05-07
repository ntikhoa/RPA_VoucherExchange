package routes

import (
	"net/http"

	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB
	productRepo       = repositories.NewProductRepo(db)
	productService    = services.NewProductService(productRepo)
	productController = controllers.NewProductController(productService)
)

func ProductRoutes(g *gin.RouterGroup, d *gorm.DB) {
	db = d

	g.POST("/", func(ctx *gin.Context) {
		providerID := productController.Create(ctx)
		if providerID == -1 {
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"provider_id": providerID,
		})
	})
}
