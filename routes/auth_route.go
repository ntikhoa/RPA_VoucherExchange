package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAuthController(db *gorm.DB) controllers.AuthController {
	services := services.NewAuthService()
	return controllers.NewAuthController(services)
}

func AuthRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initAuthController(db)

	g.POST("/register",
		func(ctx *gin.Context) {
			controller.Register(ctx)
		})
}
