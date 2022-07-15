package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initTestVoucher(db *gorm.DB) controllers.TestVoucher {
	repo := repositories.NewVoucherRepo(db)
	evService := services.NewExchangeVoucherService(repo)
	return controllers.NewTestVoucher(evService)
}

func TestVoucher(g *gin.RouterGroup, db *gorm.DB) {

	controller := initTestVoucher(db)

	g.POST("",
		middlewares.ValidateViewExchangeVoucher(),
		middlewares.ValidateTestExchangeVoucher(),
		func(ctx *gin.Context) {
			controller.TestVoucher(ctx)
		})
}
