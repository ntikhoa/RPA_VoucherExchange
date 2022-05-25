package middlewares

import (
	"log"
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
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.VOUCHER_DTO_KEY, voucherDTO)
		ctx.Next()
	}
}

func ValidatePublishedRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		publishedDTO := dto.PublishedDTO{}
		err := ctx.ShouldBind(&publishedDTO)
		if err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
		}
		ctx.Set(configs.PUBLISHED_DTO_KEY, publishedDTO)
		ctx.Next()
	}
}
