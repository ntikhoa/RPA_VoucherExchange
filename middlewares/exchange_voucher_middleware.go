package middlewares

import (
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidateViewExchangeVoucher() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		viewExchangeVoucherDTO := dto.ViewExchangeVoucherDTO{}

		ctx.Request.ParseForm()

		if len(ctx.Request.PostForm) == 0 {
			//JSON parsing
			err := ctx.ShouldBindBodyWith(&viewExchangeVoucherDTO, binding.JSON)
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

			prices, err := getUintArrayType(ctx.Request.PostForm, "prices")
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
			err := ctx.ShouldBindBodyWith(&payload, binding.JSON)
			if err != nil {
				log.Println(err)
				ctx.AbortWithError(http.StatusBadRequest, err)
				return
			}
		} else {
			//FormUrlEncoded parsing
			ids, err := getUintArrayType(ctx.Request.PostForm, "ids")
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

		log.Println("INPUT BEGIN")
		for key, values := range ctx.Request.MultipartForm.Value {
			log.Println(key, values)
		}
		log.Println("INPUT END")

		if ctx.Request.MultipartForm == nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("cannot parse form data"))
			return
		}

		products, ok := ctx.Request.MultipartForm.Value["products"]
		if !ok {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"products\" required"))
			return
		}

		log.Println(products)
		log.Println(ctx.Request.MultipartForm.Value["prices"])
		prices, err := getUintArrayType(ctx.Request.MultipartForm.Value, "prices")
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

		voucherID, err := getUintArrayType(ctx.Request.MultipartForm.Value, "voucher_id")
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

func isImageType(mimeType string) bool {
	r, _ := regexp.Compile("image/.+")
	return r.MatchString(mimeType)
}
