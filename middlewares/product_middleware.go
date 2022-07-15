package middlewares

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidateProductRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productDTO := dto.ProductDTO{}
		if err := ctx.ShouldBind(&productDTO); err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.PRODUCT_DTO_KEY, productDTO)
		ctx.Next()
	}
}

func ValidateDeleteProductsRequest() gin.HandlerFunc {
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
