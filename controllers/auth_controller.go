package controllers

import (
	"log"
	"net/http"

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
	registerDTO := dto.RegisterDTO{}
	if err := ctx.ShouldBind(&registerDTO); err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

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

}
