package middlewares

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/constants"
	"github.com/gin-gonic/gin"
)

type censor struct {
	IsApproved *bool `json:"censor" form:"censor" binding:"required"`
}

func ValidateCensorRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var censor censor
		if err := ctx.ShouldBind(&censor); err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.Set(configs.IS_APPROVED_KEY, *censor.IsApproved)
		ctx.Next()
	}
}

func ValidateSearchByDateQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		values := ctx.Request.URL.Query()
		fromDate_string := values["from_date"]
		toDate_string := values["to_date"]
		if len(fromDate_string) != 1 || len(toDate_string) != 1 {
			ctx.AbortWithError(http.StatusBadRequest, errors.New(`'from_date' and 'to_date' required`))
			return
		}

		fromDate, err := time.Parse(constants.SEARCH_DATE_FORMAT, fromDate_string[0])
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		toDate, err := time.Parse(constants.SEARCH_DATE_FORMAT, toDate_string[0])
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		ctx.Set(configs.FROM_DATE, fromDate)
		ctx.Set(configs.TO_DATE, toDate)
		ctx.Next()
	}
}
