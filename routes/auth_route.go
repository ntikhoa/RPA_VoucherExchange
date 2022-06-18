package routes

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(g *gin.RouterGroup, db *gorm.DB) {
	controller := initAuthController(db)

	g.POST("/login",
		middlewares.ValidateLoginRequest(),
		func(ctx *gin.Context) {
			controller.Login(ctx, constants.ROLE_SALE)
		})

	g.POST("/auto_login",
		middlewares.AuthorizeJwt(db),
		func(ctx *gin.Context) {
			controller.AutoLogin(ctx)
		})
}
