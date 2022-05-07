package middlewares

import (
	"github.com/RPA_VoucherExchange/constants"
	"github.com/gin-gonic/gin"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(constants.PROVIDER_ID_KEY, 1)
	}
}
