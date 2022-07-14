package middlewares

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
)

func ValidateViewExchangeVoucher() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		viewExchangeVoucherDTO := dto.ViewExchangeVoucherDTO{}

		ctx.Request.ParseForm()

		if len(ctx.Request.PostForm) == 0 {
			//JSON parsing
			err := ctx.ShouldBind(&viewExchangeVoucherDTO)
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
		} else {
			//FormUrlEncoded parsing
			products, ok := ctx.Request.PostForm["products"]
			if !ok {
				ctx.AbortWithError(http.StatusBadRequest, errors.New("\"products\" required"))
				return
			}

			prices, err := getUintArrayTypeFormURL(ctx, "prices")
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}

			viewExchangeVoucherDTO = dto.ViewExchangeVoucherDTO{
				Products: products,
				Prices:   prices,
			}
		}

		if len(viewExchangeVoucherDTO.Products) != len(viewExchangeVoucherDTO.Prices) {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("the number of products and prices do not match"))
		}

		ctx.Set(configs.VIEW_EXCHANGE_VOUCHER_DTO_KEY, viewExchangeVoucherDTO)
		ctx.Next()
	}
}

func ValidateTestExchangeVoucher() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		viewExchangeVoucherDTO := ctx.MustGet(configs.VIEW_EXCHANGE_VOUCHER_DTO_KEY).(dto.ViewExchangeVoucherDTO)

		ctx.Request.ParseForm()

		var payload dto.PayLoad

		if len(ctx.Request.PostForm) == 0 {
			//JSON parsing
			err := ctx.ShouldBind(&payload)
			if err != nil {
				log.Println(err)
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
		} else {
			//FormUrlEncoded parsing
			ids, err := getUintArrayTypeFormURL(ctx, "ids")
			if err != nil {
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
			payload.IDs = ids
		}

		testVoucherDTO := dto.TestExchangeVoucherDTO{
			ViewExchangeVoucherDTO: viewExchangeVoucherDTO,
			VoucherIDs:             payload.IDs,
		}

		ctx.Set(configs.TEST_EXCHANGE_VOUCHER_DTO_KEY, testVoucherDTO)
		ctx.Next()
	}
}

func ValidateExchangeVoucher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.ParseMultipartForm(32 << 20) // 32MB + 10MB

		if ctx.Request.MultipartForm == nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse form data"))
			return
		}

		products, ok := ctx.Request.MultipartForm.Value["products"]
		if !ok {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"products\" required"))
			return
		}

		prices, err := getUintArrayType(ctx, "prices")
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if len(products) != len(prices) {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("the number of products and prices do not match"))
		}

		files := ctx.Request.MultipartForm.File["files"]
		if len(files) == 0 {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"files\" required"))
			return
		}
		for _, file := range files {
			mimeType := file.Header["Content-Type"][0]
			if !isImageType(mimeType) {
				ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid file type"))
				return
			}
		}

		voucherID, err := getUintArrayType(ctx, "voucher_id")
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		customerName, ok := ctx.Request.MultipartForm.Value["customer_name"]
		if !ok || len(customerName) < 1 {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"customer_name\" required"))
			return
		}

		customerPhone, ok := ctx.Request.MultipartForm.Value["customer_phone"]
		if !ok || len(customerPhone) < 1 {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"customer_phone\" required"))
			return
		}

		transactionID, ok := ctx.Request.MultipartForm.Value["transaction_id"]
		if !ok || len(customerPhone) < 1 {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"transaction_id\" required"))
			return
		}

		exchangeVoucherDTO := dto.ExchangeVoucherDTO{
			ViewExchangeVoucherDTO: dto.ViewExchangeVoucherDTO{
				Products: products,
				Prices:   prices,
			},
			TransactionID: transactionID[0],
			VoucherID:     voucherID[0],
			CustomerName:  customerName[0],
			CustomerPhone: customerPhone[0],
		}

		ctx.Set(configs.RECEIPT_IMAGE_FILES_KEY, files)
		ctx.Set(configs.EXCHANGE_VOUCHER_DTO, exchangeVoucherDTO)
		ctx.Next()
	}
}

func getUintArrayTypeFormURL(ctx *gin.Context, key string) ([]uint, error) {
	pricesStr, ok := ctx.Request.PostForm[key]
	if !ok {
		return nil, errors.New("\"" + key + "\"" + " required")
	}
	var prices []uint
	for _, priceStr := range pricesStr {
		price, err := strconv.ParseUint(priceStr, 10, 64)
		if err != nil {
			return prices, errors.New("\"" + key + "\" " + " invalid type")
		}
		prices = append(prices, uint(price))
	}

	return prices, nil
}

func getUintArrayType(ctx *gin.Context, key string) ([]uint, error) {
	pricesStr, ok := ctx.Request.MultipartForm.Value[key]
	if !ok {
		return nil, errors.New("\"" + key + "\"" + " required")
	}
	var prices []uint
	for _, priceStr := range pricesStr {
		price, err := strconv.ParseUint(priceStr, 10, 64)
		if err != nil {
			return prices, errors.New("\"" + key + "\" " + " invalid type")
		}
		prices = append(prices, uint(price))
	}

	return prices, nil
}

func isImageType(mimeType string) bool {
	r, _ := regexp.Compile("image/.+")
	return r.MatchString(mimeType)
}
