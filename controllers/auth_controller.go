package controllers

import (
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
	s services.AuthService
}

func NewAuthController(s services.AuthService) AuthController {
	return &authController{
		s: s,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	registerDTO := dto.RegisterDTO{}
	if err := ctx.ShouldBind(&registerDTO); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": registerDTO,
	})
}

func (c *authController) Login(ctx *gin.Context) {

}
