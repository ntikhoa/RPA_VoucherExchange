package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

func GetSearchFromURL() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		values := ctx.Request.URL.Query()
		query := values["query"]
		if len(query) != 1 {
			err := errors.New("search query is required")
			log.Println(err)
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Set(configs.SEARCH_QUERY_KEY, query[0])
		ctx.Next()
	}
}
