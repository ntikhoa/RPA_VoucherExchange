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
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) AuthController {
	return &authController{
		service: service,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	registerDTO := ctx.MustGet(configs.REGISTER_KEY).(dto.RegisterDTO)
	err := c.service.Register(registerDTO)
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

func (c *authController) Login(ctx *gin.Context) {
	loginDTO := ctx.MustGet(configs.LOGIN_KEY).(dto.LoginDTO)

	token, err := c.service.Login(loginDTO)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"status": 200,
		"data": gin.H{
			"token": token,
		},
		"error":   nil,
		"message": "register successfully",
	})
}
