package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

func GetPageFromURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value := ctx.Request.URL.Query()
		pageQuery := value["page"]
		if len(pageQuery) != 1 {
			err := errors.New("page query is required")
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		pageConv, err := strconv.ParseInt(pageQuery[0], 10, 64)
		if err != nil || pageConv <= 0 {
			err := errors.New("invalid page query, page should be greater than or equal to 1")
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.PAGE_QUERY_KEY, int(pageConv))
		ctx.Next()
	}
}
