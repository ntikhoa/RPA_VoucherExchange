package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

func GetProductNameFromURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		values := ctx.Request.URL.Query()
		productNameQuery := values["product_name"]
		if len(productNameQuery) != 1 {
			err := errors.New("Invalid Product Name argument.")
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		productName := string(productNameQuery[0])
		ctx.Set(configs.PRODUCT_NAME_KEY, productName)
		ctx.Next()
	}
}
