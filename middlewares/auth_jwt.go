package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if len(authHeader) <= len(BEARER_SCHEMA) {
			ctx.AbortWithError(http.StatusUnauthorized, errors.New(constants.INVALID_TOKEN_ERROR))
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := services.NewJWTService().ValidateToken(tokenString)
		if err != nil {
			log.Println(err)
			ctx.AbortWithError(http.StatusUnauthorized, errors.New(constants.INVALID_TOKEN_ERROR))
			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[EmployeeID]: ", claims["employee_id"])
			log.Println("Claims[ProviderID]: ", claims["provider_id"])
			log.Println("Claims[Issuer]: ", claims["iss"])
			log.Println("Claims[IssuedAt]: ", claims["iat"])
			log.Println("Claims[ExpiresAt]: ", claims["exp"])

			providerIDStr := claims["provider_id"].(string)
			providerID, err := strconv.ParseUint(providerIDStr, 10, 64)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			ctx.Set(configs.TOKEN_PROVIDER_ID_KEY, uint(providerID))

			employeeIDStr := claims["employee_id"].(string)
			employeeID, err := strconv.ParseUint(employeeIDStr, 10, 64)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			ctx.Set(configs.TOKEN_EMPLOYEE_ID_KEY, uint(employeeID))

			ctx.Next()
		} else {
			log.Println(err)
			ctx.AbortWithError(http.StatusUnauthorized, errors.New(constants.INVALID_TOKEN_ERROR))
			return
		}
	}
}
