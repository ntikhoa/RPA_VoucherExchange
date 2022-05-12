package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

//get the id from route routes/:id
func GetIDFromURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		productID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			log.Println(err)
			err := errors.New("invalid ID param")
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.ID_PARAM_KEY, uint(productID))
		ctx.Next()
	}
}
