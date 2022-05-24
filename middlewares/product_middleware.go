package middlewares

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
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
