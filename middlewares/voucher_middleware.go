package middlewares

import (
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
)

func ValidateVoucherRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		voucherDTO := dto.VoucherDTO{}
		err := ctx.ShouldBind(&voucherDTO)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.Set(configs.VOUCHER_DTO_KEY, voucherDTO)
		ctx.Next()
	}
}
