package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/constants"
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

func ValidateRegisterSaleRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
		registerSaleDTO := dto.RegisterSaleDTO{}
		if err := ctx.ShouldBind(&registerSaleDTO); err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		registerDTO := registerSaleDTO.ToRegisterDTO(providerID)
		ctx.Set(configs.REGISTER_KEY, registerDTO)
		ctx.Next()
	}
}

func AuthorizeAdminRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleID := ctx.MustGet(configs.ACCOUNT_ROLE_ID_KEY).(uint)
		if roleID != constants.ROLE_ADMIN {
			ctx.AbortWithError(http.StatusForbidden, errors.New(constants.AUTHORIZE_ERROR))
			return
		}

		ctx.Next()
	}
}
