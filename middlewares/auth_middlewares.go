package middlewares

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
)

func ValidateRegisterRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		registerDTO := dto.RegisterDTO{}
		if err := ctx.ShouldBind(&registerDTO); err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.REGISTER_KEY, registerDTO)
		ctx.Next()
	}
}

func ValidateLoginRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		loginDTO := dto.LoginDTO{}
		err := ctx.ShouldBind(&loginDTO)
		if err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set("login_key", loginDTO)
		ctx.Next()
	}
}
