package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if !ctx.IsAborted() {
			return
		}

		status := ctx.Writer.Status()
		err := http.StatusText(status)
		msg := "No message available"

		if len(ctx.Errors) > 0 {
			msg = ctx.Errors.Last().Err.Error()
		}

		ctx.JSON(status, gin.H{
			"status":  status,
			"data":    nil,
			"error":   err,
			"message": msg,
		})
	}
}
