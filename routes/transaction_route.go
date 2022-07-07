package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initTransactionController(db *gorm.DB) controllers.TransactionController {
	repo := repositories.NewTransactionRepo(db)
	service := services.NewTransactionService(repo)
	return controllers.NewTransactionController(service)
}

func TransactionRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initTransactionController(db)

	g.GET("", func(ctx *gin.Context) {
		controller.FindAll(ctx)
	})
}
