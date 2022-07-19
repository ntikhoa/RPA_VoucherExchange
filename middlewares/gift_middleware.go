package middlewares

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidateGiftRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		giftDTO := dto.GiftDTO{}
		if err := ctx.ShouldBind(&giftDTO); err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.GIFT_DTO_KEY, giftDTO)
		ctx.Next()
	}
}

func ValidateDeleteGiftsRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Request.ParseForm()

		var payload dto.PayLoad

		if len(ctx.Request.PostForm) == 0 {
			//JSON parsing
			err := ctx.ShouldBindBodyWith(&payload, binding.JSON)
			if err != nil {
				log.Println(err)
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
		} else {
			//FormUrlEncoded parsing
			ids, err := getUintArrayType(ctx.Request.PostForm, "ids")
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
			payload.IDs = ids
		}
		ctx.Set(configs.PAYLOAD_DTO_KEY, payload.IDs)
		ctx.Next()
	}
}
