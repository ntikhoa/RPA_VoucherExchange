package middlewares

import (
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
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
