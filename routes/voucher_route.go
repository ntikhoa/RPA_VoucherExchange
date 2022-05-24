package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initVoucherController(db *gorm.DB) controllers.VoucherController {
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	voucherRepo := repositories.NewVoucherRepo(db)
	voucherService := services.NewVoucherService(voucherRepo)
	return controllers.NewVoucherController(voucherService, productService)
}

func VoucherRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initVoucherController(db)

	g.POST("",
		func(ctx *gin.Context) {
			controller.Create(ctx)
		})

	g.GET(":id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			controller.FindByID(ctx)
		})
}
