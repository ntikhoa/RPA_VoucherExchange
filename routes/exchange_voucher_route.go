package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initExchangeVoucherController(db *gorm.DB) controllers.ExchangeVoucherController {
	voucherRepo := repositories.NewVoucherRepo(db)
	receiptRepo := repositories.NewReceiptRepo(db)

	evService := services.NewExchangeVoucherService(voucherRepo)
	receiptService := services.NewReceiptService(receiptRepo)
	imageService := services.NewImageService()
	return controllers.NewExchangeVoucherController(evService, imageService, receiptService)
}

func ExchangeVoucherRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initExchangeVoucherController(db)

	g.POST("/view",
		middlewares.ValidateViewExchangeVoucher(),
		func(ctx *gin.Context) {
			controller.ViewExchangeVoucher(ctx)
		})

	g.POST("",
		middlewares.ValidateExchangeVoucher(),
		func(ctx *gin.Context) {
			controller.ExchangeVoucher(ctx)
		})

}
