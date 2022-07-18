package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initVoucherController(db *gorm.DB) controllers.VoucherController {
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	voucherRepo := repositories.NewVoucherRepo(db)
	voucherService := services.NewVoucherService(voucherRepo)
	giftRepo := repositories.NewGiftRepo(db)
	giftService := services.NewGiftService(giftRepo)
	return controllers.NewVoucherController(voucherService, productService, giftService)
}

func VoucherRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initVoucherController(db)

	g.POST("",
		middlewares.ValidateVoucherRequest(),
		func(ctx *gin.Context) {
			controller.Create(ctx)
		})

	g.GET(":id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			controller.FindByID(ctx)
		})

	g.GET("",
		middlewares.GetPageFromURL(),
		func(ctx *gin.Context) {
			controller.FindAll(ctx)
		})

	g.GET("/search",
		middlewares.GetSearchFromURL(),
		func(ctx *gin.Context) {
			controller.Search(ctx)
		})

	g.DELETE(":id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			controller.Delete(ctx)
		})

	g.POST("publish",
		middlewares.ValidatePublishedRequest(),
		func(ctx *gin.Context) {
			controller.Publish(ctx)
		})

	g.PUT(":id",
		middlewares.GetIDFromURL(),
		middlewares.ValidateVoucherRequest(),
		func(ctx *gin.Context) {
			controller.Update(ctx)
		})
}
