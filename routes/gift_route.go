package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initGiftController(db *gorm.DB) controllers.GiftController {
	giftRepo := repositories.NewGiftRepo(db)
	giftService := services.NewGiftService(giftRepo)
	return controllers.NewGiftController(giftService)
}

func GiftRoutes(g *gin.RouterGroup, db *gorm.DB) {
	giftController := initGiftController(db)

	g.POST("",
		middlewares.ValidateGiftRequest(),
		func(ctx *gin.Context) {
			giftController.CreateGift(ctx)
		})

	g.DELETE("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			giftController.DeleteGift(ctx)
		})

	g.PUT("/:id",
		middlewares.GetIDFromURL(),
		middlewares.ValidateGiftRequest(),
		func(ctx *gin.Context) {
			giftController.UpdateGift(ctx)
		})

	g.GET("", func(ctx *gin.Context) {
		giftController.FindAllGift(ctx)
	})

	g.GET("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			giftController.FindGiftByID(ctx)
		})
}
