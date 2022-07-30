package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initTransactionSaleController(db *gorm.DB) controllers.TransactionSaleController {
	repo := repositories.NewReceiptRepo(db)
	service := services.NewReceiptService(repo)
	return controllers.NewTransactionSaleController(service)
}

func TransactionSaleRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initTransactionSaleController(db)

	g.GET("",
		middlewares.GetPageFromURL(),
		func(ctx *gin.Context) {
			controller.FindAll(ctx)
		})

	g.GET("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			controller.FindByID(ctx)
		})
	g.GET("search",
		middlewares.ValidateSearchByDateQuery(),
		func(ctx *gin.Context) {
			controller.FindBetweenDates(ctx)
		})
}
