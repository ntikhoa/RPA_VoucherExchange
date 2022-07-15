package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAccountController(db *gorm.DB) controllers.AccountController {
	accountRepo := repositories.NewAccountRepo(db)
	accountService := services.NewAccountService(accountRepo)
	return controllers.NewAccountController(accountService)
}

func AccountRoutes(g *gin.RouterGroup, db *gorm.DB) {

	controller := initAccountController(db)

	g.GET("",
		middlewares.GetPageFromURL(),
		func(ctx *gin.Context) {
			controller.GetAccount(ctx)
		})
	g.GET("/search",
		middlewares.GetSearchFromURL(),
		func(ctx *gin.Context) {
			controller.Search(ctx)
		})
}
