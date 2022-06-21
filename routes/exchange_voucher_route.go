package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/gin-gonic/gin"
)

func ExchangeVoucherRoutes(g *gin.RouterGroup) {

	controller := controllers.NewExchangeVoucherController()

	g.POST("",
		middlewares.ValidateExchangeVoucher(),
		func(ctx *gin.Context) {
			controller.ExchangeVoucher(ctx)
		})
}
