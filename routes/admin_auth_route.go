package routes

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAuthController(db *gorm.DB) controllers.AuthController {
	providerRepo := repositories.NewProviderRepo(db)
	accountRepo := repositories.NewAccountRepo(db)
	authService := services.NewAuthService(accountRepo, providerRepo)
	jwtService := services.NewJWTService()
	return controllers.NewAuthController(authService, jwtService)
}

func AdminAuthRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initAuthController(db)

	g.POST("/register_admin",
		middlewares.ValidateRegisterRequest(),
		func(ctx *gin.Context) {
			controller.Register(ctx, constants.ROLE_ADMIN)
		})

	g.POST("/login",
		middlewares.ValidateLoginRequest(),
		func(ctx *gin.Context) {
			controller.Login(ctx, constants.ROLE_ADMIN)
		})

	g.POST("/register_sales",
		middlewares.AuthorizeJwt(db),
		middlewares.AuthorizeAdminRole(),
		middlewares.ValidateRegisterSaleRequest(),
		func(ctx *gin.Context) {
			controller.Register(ctx, constants.ROLE_SALE)
		})
}
