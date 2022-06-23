package middlewares

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

func ValidateExchangeVoucher() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Request.ParseMultipartForm(32 << 20) // 32MB + 10MB

		files := ctx.Request.MultipartForm.File["files"]

		for _, file := range files {
			mimeType := file.Header["Content-Type"][0]
			if !isImageType(mimeType) {
				ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid file type"))
				return
			}
		}

		ocrProducts, ok := ctx.Request.MultipartForm.Value["ocr_products"]
		if !ok {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("\"ocr_products\" required"))
			return
		}

		ocrPrices, err := getUintArrayType(ctx, "ocr_prices")
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
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

		ctx.Set(configs.RECEIPT_IMAGE_FILES_KEY, files)
		ctx.Set(configs.ORIGINAL_PRODUCTS_KEY, ocrProducts)
		ctx.Set(configs.ORIGINAL_PRICES_KEY, ocrPrices)
		ctx.Set(configs.EDITED_PRODUCTS_KEY, products)
		ctx.Set(configs.EDITED_PRICES_KEY, prices)

		ctx.Next()
	}
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
			return nil, errors.New("\"" + key + "\" " + " invalid type")
		}
		prices = append(prices, uint(price))
	}

	return prices, nil
}

func isImageType(mimeType string) bool {
	r, _ := regexp.Compile("image/.+")
	return r.MatchString(mimeType)
}
