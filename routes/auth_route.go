package routes

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(g *gin.RouterGroup, db *gorm.DB) {
	controller := initAuthController(db)

	g.POST("/login", func(ctx *gin.Context) {
		controller.Login(ctx, constants.ROLE_SALE)
	})
}
