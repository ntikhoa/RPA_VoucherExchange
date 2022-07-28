package middlewares

import (
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
)

func ValidateTransferRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dto := dto.CreateTransferGiftsDTO{}
		err := ctx.ShouldBind(&dto)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.Set(configs.TRANSFER_DTO_KEY, dto)
		ctx.Next()
	}
}
