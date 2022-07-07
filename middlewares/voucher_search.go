package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

func GetVoucherNameFromUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		values := ctx.Request.URL.Query()
		voucherNameQuery := values["voucher_name"]
		if len(voucherNameQuery) != 1 {
			err := errors.New("Invalid Voucher Name argument.")
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		//Just to be sure it is string
		voucherName := string(voucherNameQuery[0])
		ctx.Set(configs.VOUCHER_NAME_KEY, voucherName)
		ctx.Next()
	}
}
