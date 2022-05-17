package routes

import (
	"github.com/RPA_VoucherExchange/controllers"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func initAuthController(db *gorm.DB) controllers.AuthController {
	providerRepo := repositories.NewProviderRepo(db)
	employeeRepo := repositories.NewEmployeeRepo(db)
	authService := services.NewAuthService(employeeRepo, providerRepo)
	jwtService := services.NewJWTService()
	return controllers.NewAuthController(authService, jwtService)
}

func AuthRoutes(g *gin.RouterGroup, db *gorm.DB) {
	controller := initAuthController(db)

	g.POST("/register",
		middlewares.ValidateRegisterRequest(),
		func(ctx *gin.Context) {
			controller.Register(ctx)
		})

	g.POST("/login",
		middlewares.ValidateLoginRequest(),
		func(ctx *gin.Context) {
			controller.Login(ctx)
		})
}
