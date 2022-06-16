package controllers

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(ctx *gin.Context, roleID uint)
	Login(ctx *gin.Context, roleID uint)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController(authService services.AuthService, jwtService services.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Register(ctx *gin.Context, roleID uint) {
	registerDTO := ctx.MustGet(configs.REGISTER_KEY).(dto.RegisterDTO)
	err := c.authService.Register(registerDTO, roleID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status":  200,
		"data":    nil,
		"error":   nil,
		"message": "register successfully",
	})
}

func (c *authController) Login(ctx *gin.Context, roleID uint) {
	loginDTO := ctx.MustGet(configs.LOGIN_KEY).(dto.LoginDTO)

	employee, err := c.authService.Login(loginDTO, roleID)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	token, err := c.jwtService.GenerateToken(employee)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"data": gin.H{
			"token": token,
			"name":  employee.Name,
		},
		"error":   nil,
		"message": "login successfully",
	})
}
