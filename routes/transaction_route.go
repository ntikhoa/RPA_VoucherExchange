package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initTransactionController(db *gorm.DB) controllers.TransactionController {
	repo := repositories.NewReceiptRepo(db)
	service := services.NewReceiptService(repo)
	return controllers.NewTransactionController(service)
}

func TransactionRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initTransactionController(db)

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
	g.PUT("/:id",
		middlewares.GetIDFromURL(),
		middlewares.ValidateCensorRequest(),
		func(ctx *gin.Context) {
			controller.Censor(ctx)
		})
}
