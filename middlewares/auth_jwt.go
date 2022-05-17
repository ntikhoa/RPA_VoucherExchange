package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Invalid Token"))
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := services.ValidateToken(tokenString)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Invalid Token"))
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[EmployeeID]: ", claims["employee_id"])
			log.Println("Claims[ProviderID]: ", claims["provider_id"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

			ctx.Set(configs.TOKEN_DATA_KEY, claims)
			ctx.Next()
		} else {
			log.Println(err)
			ctx.AbortWithError(http.StatusUnauthorized, errors.New("Invalid Token"))
			return
		}
	}
}
