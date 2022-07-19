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
			giftController.Create(ctx)
		})

	g.DELETE("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			giftController.Delete(ctx)
		})

	g.PUT("/:id",
		middlewares.GetIDFromURL(),
		middlewares.ValidateGiftRequest(),
		func(ctx *gin.Context) {
			giftController.Update(ctx)
		})

	g.GET("",
		middlewares.GetPageFromURL(),
		func(ctx *gin.Context) {
			giftController.FindAll(ctx)
		})

	g.GET("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			giftController.FindByID(ctx)
		})

	g.GET("/search",
		middlewares.GetSearchFromURL(),
		func(ctx *gin.Context) {
			giftController.Search(ctx)
		})

	g.GET("/getAll",
		func(ctx *gin.Context) {
			giftController.GetAll(ctx)
		})

	g.POST("/delete_gifts",
		middlewares.ValidateDeleteGiftsRequest(),
		func(ctx *gin.Context) {
			giftController.DeleteGifts(ctx)
		})
}
