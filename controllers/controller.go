package controllers

import (
	"net/http"

	custom_error "github.com/RPA_VoucherExchange/custom_error"
	"github.com/gin-gonic/gin"
)

//check for custom error and abort with the coresponded status code
//use for checking custom message sent from the services layer
func abortCustomError(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	switch err.(type) {
	case *custom_error.BadRequestError:
		status = http.StatusBadRequest

	case *custom_error.ForbiddenError:
		status = http.StatusForbidden

	case *custom_error.NotFoundError:
		status = http.StatusNotFound

	case *custom_error.ConflictError:
		status = http.StatusConflict
	}
	ctx.AbortWithError(status, err)
}
