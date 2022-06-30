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
	repo := repositories.NewVoucherRepo(db)
	service := services.NewExchangeVoucherService(repo)
	return controllers.NewExchangeVoucherController(service)
}

func ExchangeVoucherRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initExchangeVoucherController(db)

	g.POST("/view",
		middlewares.AuthorizeJwt(db),
		middlewares.ValidateViewExchangeVoucher(),
		func(ctx *gin.Context) {
			controller.ViewExchangeVoucher(ctx)
		})

	g.POST("",
		middlewares.AuthorizeJwt(db),
		middlewares.ValidateViewExchangeVoucher(),
		middlewares.ValidateExchangeVoucher(),
		func(ctx *gin.Context) {
			controller.ExchangeVoucher(ctx)
		})

}
