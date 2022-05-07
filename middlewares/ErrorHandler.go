package middlewares

import (
	"log"
	"net/http"
	"time"

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

		for _, errMsg := range ctx.Errors {
			log.Println(errMsg.Err)
			msg = errMsg.Err.Error()
		}

		ctx.JSON(status, gin.H{
			"timestamp": time.Now(),
			"status":    status,
			"error":     err,
			"message":   msg,
			"path":      ctx.Request.URL.Path,
		})
	}
}
