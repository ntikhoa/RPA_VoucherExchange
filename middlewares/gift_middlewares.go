package middlewares

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
)

func ValidateGiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		giftDTO := dto.GiftDTO{}
		if err := ctx.ShouldBind(&giftDTO); err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		gift := giftDTO.ToEntity()
		ctx.Set(configs.GIFT_KEY, gift)
		ctx.Next()
	}
}
