package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAccountSaleController(db *gorm.DB) controllers.AccountSaleController {
	accountRepo := repositories.NewAccountRepo(db)
	accountService := services.NewAccountService(accountRepo)
	return controllers.NewAccountSaleController(accountService)
}

func AccountSaleRoutes(g *gin.RouterGroup, db *gorm.DB) {

	controller := initAccountSaleController(db)

	g.GET("",
		func(ctx *gin.Context) {
			controller.GetAccountProfile(ctx)
		})
}
