package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initVoucherController(db *gorm.DB) controllers.VoucherController {
	db.Unscoped().Model(&entities.Voucher{}).Where("id = ?", 2).Update("deleted_at", nil)
	db.Unscoped().Model(&entities.Gift{}).Where("voucher_id = ?", 2).Update("deleted_at", nil)
	db.Unscoped().Model(&entities.VoucherProduct{}).Where("voucher_id = ?", 2).Update("deleted_at", nil)

	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	voucherRepo := repositories.NewVoucherRepo(db)
	voucherService := services.NewVoucherService(voucherRepo)
	return controllers.NewVoucherController(voucherService, productService)
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
