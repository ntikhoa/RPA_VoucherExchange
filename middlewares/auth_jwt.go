package middlewares

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/repositories"
	"github.com/RPA_VoucherExchange/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func AuthorizeJwt(db *gorm.DB) gin.HandlerFunc {
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

			providerIDStr := claims["provider_id"].(string)
			providerIDuint64, err := strconv.ParseUint(providerIDStr, 10, 64)
			providerID := uint(providerIDuint64)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			employeeIDStr := claims["employee_id"].(string)
			employeeIDuint64, err := strconv.ParseUint(employeeIDStr, 10, 64)
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			employeeID := uint(employeeIDuint64)

			employee, err := repositories.NewEmployeeRepo(db).FindEmployee(employeeID, providerID)
			if err != nil {
				ctx.AbortWithError(http.StatusUnauthorized, errors.New(constants.INVALID_TOKEN_ERROR))
				return
			}

			if !(employee.IssueAt.Format(constants.JWT_DATE_FORMAT) == claims["issue_at"].(string)) {
				ctx.AbortWithError(http.StatusUnauthorized, errors.New(constants.INVALID_TOKEN_ERROR))
				return
			}

			ctx.Set(configs.TOKEN_PROVIDER_ID_KEY, providerID)
			ctx.Set(configs.TOKEN_EMPLOYEE_ID_KEY, employeeID)
			ctx.Next()
		} else {
			log.Println(err)
			ctx.AbortWithError(http.StatusUnauthorized, errors.New(constants.INVALID_TOKEN_ERROR))
			return
		}
	}
}
