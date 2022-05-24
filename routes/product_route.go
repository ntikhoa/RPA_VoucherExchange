package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initProductController(db *gorm.DB) controllers.ProductController {
	productRepo := repositories.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	return controllers.NewProductController(productService)
}

func ProductRoutes(g *gin.RouterGroup, db *gorm.DB) {
	productController := initProductController(db)

	g.POST("",
		middlewares.ValidateProductRequest(),
		func(ctx *gin.Context) {
			productController.Create(ctx)
		})

	g.DELETE("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			productController.Delete(ctx)
		})

	g.PUT("/:id",
		middlewares.GetIDFromURL(),
		middlewares.ValidateProductRequest(),
		func(ctx *gin.Context) {
			productController.Update(ctx)
		})

	g.GET("",
		middlewares.GetPageFromURL(),
		func(ctx *gin.Context) {
			productController.FindAll(ctx)
		})

	g.GET("/:id",
		middlewares.GetIDFromURL(),
		func(ctx *gin.Context) {
			productController.FindByID(ctx)
		})
}
