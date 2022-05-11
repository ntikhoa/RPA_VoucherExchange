package middlewares

import "github.com/gin-gonic/gin"

func SetContentTypeJSON() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
