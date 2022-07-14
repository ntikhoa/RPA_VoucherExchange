package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strconv"

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

func ValidateDeleteProductsRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productsID []uint
		ctx.Request.ParseForm()
		productIDsStr, ok := ctx.Request.Form["ids"]
		if !ok {
			var payload dto.PayLoad
			if err := ctx.ShouldBind(&payload); err != nil {
				log.Println(err)
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
			productsID = payload.IDs
		} else {
			for _, idStr := range productIDsStr {
				productIDuint64, err := strconv.ParseUint(idStr, 10, 64)
				if err != nil {
					log.Println(err)
					ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid request body"))
					return
				}
				productsID = append(productsID, uint(productIDuint64))
			}
		}

		ctx.Set(configs.PAYLOAD_DTO_KEY, productsID)
		ctx.Next()
	}
}
