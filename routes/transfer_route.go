package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initTransferController(db *gorm.DB) controllers.TransferController {

	giftRepo := repositories.NewGiftRepo(db)
	accountRepo := repositories.NewAccountRepo(db)
	transferRepo := repositories.NewTransferRepo(db)

	giftService := services.NewGiftService(giftRepo)
	accountService := services.NewAccountService(accountRepo)
	transferService := services.NewTransferService(transferRepo)

	return controllers.NewTransferController(giftService, accountService, transferService)
}

func TransferRoutes(g *gin.RouterGroup, db *gorm.DB) {

	controller := initTransferController(db)

	g.POST("",
		middlewares.ValidateTransferRequest(),
		func(ctx *gin.Context) {
			controller.TransferGift(ctx)
		})

	//account id
	g.GET("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			controller.GetTransferGift(ctx)
		})
}
