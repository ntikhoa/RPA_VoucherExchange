package controllers

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/services"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"github.com/gin-gonic/gin"
)

type ExchangeVoucherController interface {
	ViewExchangeVoucher(ctx *gin.Context)
	ExchangeVoucher(ctx *gin.Context)
}

type exchangeVoucherController struct {
	evService      services.ExchangeVoucherService
	imageService   services.ImageService
	receiptService services.ReceiptService
}

func NewExchangeVoucherController(evService services.ExchangeVoucherService,
	imageService services.ImageService,
	receiptService services.ReceiptService) ExchangeVoucherController {
	return &exchangeVoucherController{
		evService:      evService,
		imageService:   imageService,
		receiptService: receiptService,
	}
}

func (c *exchangeVoucherController) ViewExchangeVoucher(ctx *gin.Context) {
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	viewVoucherExchangeDTO := ctx.MustGet(configs.VIEW_EXCHANGE_VOUCHER_DTO_KEY).(dto.ViewExchangeVoucherDTO)

	vouchers, err := c.evService.ViewExchangeVoucher(providerID, viewVoucherExchangeDTO)
	if err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	res := viewmodel.NewExchangeVoucherListResponse(vouchers)

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": gin.H{
			"vouchers": res,
		},
		"error":   nil,
		"message": "Vouchers found successfully.",
	})
}

func (c *exchangeVoucherController) ExchangeVoucher(ctx *gin.Context) {
	accountID := ctx.MustGet(configs.TOKEN_ACCOUNT_ID_KEY).(uint)
	providerID := ctx.MustGet(configs.TOKEN_PROVIDER_ID_KEY).(uint)
	files := ctx.MustGet(configs.RECEIPT_IMAGE_FILES_KEY).([]*multipart.FileHeader)
	dto := ctx.MustGet(configs.EXCHANGE_VOUCHER_DTO).(dto.ExchangeVoucherDTO)

	if err := c.evService.GetVoucherExchange(providerID, dto.ViewExchangeVoucherDTO, dto.VoucherID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	filesName, err := c.imageService.UploadObjects(files)
	if err != nil {
		log.Println(err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := c.receiptService.Create(dto, filesName, accountID); err != nil {
		log.Println(err)
		abortCustomError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"data":    nil,
		"error":   nil,
		"message": "Vouchers exchange successfully.",
	})
}
