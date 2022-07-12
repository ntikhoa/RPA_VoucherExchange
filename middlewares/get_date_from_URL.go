package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

func GetDatesFromURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		values := ctx.Request.URL.Query()
		fromDate_strings := values["from_date"]
		toDate_strings := values["to_date"]
		if len(fromDate_strings) != 1 || len(toDate_strings) != 1 {
			err := errors.New("Invalid dates.")
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.FROM_DATE, fromDate_strings[0])
		ctx.Set(configs.TO_DATE, toDate_strings[0])
		ctx.Next()
	}

}
