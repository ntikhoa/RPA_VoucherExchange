package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ExchangeVoucherRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := controllers.NewExchangeVoucherController()

	g.POST("",
		middlewares.ValidateExchangeVoucher(),
		func(ctx *gin.Context) {
			repo := repositories.NewVoucherRepo(db)
			vouchers, err := repo.FindVoucherExchange(1)
			if err != nil {
				panic(err)
			}
			ctx.JSON(200, gin.H{
				"data": vouchers,
			})
			return
			controller.ExchangeVoucher(ctx)

		})
}
