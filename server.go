package main

import (
	"net/http"
	"os"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/routes"
	"github.com/gin-gonic/gin"
)

var (
	conn = configs.NewDBConnection()
)

func main() {
	// for loading local .env file
	// utils.LoadDotEnv()

	port := os.Getenv("PORT")
	server := gin.New()

	conn.ConnectDB()
	defer conn.CloseDB()
	// conn.Init()
	db := conn.GetDB()
	// utils.SeedProducts(db)
	// utils.SeedVouchers(db, 1)

	server.Use(middlewares.SetHeader())
	server.Use(middlewares.ErrorHandler())

	server.POST("/echo_post", func(ctx *gin.Context) {
		ctx.Request.ParseForm()

		ctx.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    ctx.Request.PostForm,
			"error":   nil,
			"message": "Vouchers found successfully.",
		})
	})

	apiRoutesV1 := server.Group("/api/v1")
	{

		apiAdminRoutes := apiRoutesV1.Group("/admin")
		{
			apiAuthRoutes := apiAdminRoutes.Group("/auth")
			{
				routes.AdminAuthRoutes(apiAuthRoutes, db)
			}

			apiProductRoutes := apiAdminRoutes.Group("/products")
			{
				apiProductRoutes.Use(middlewares.AuthorizeJwt(db))
				apiProductRoutes.Use(middlewares.AuthorizeAdminRole())
				routes.ProductRoutes(apiProductRoutes, db)
			}

			apiGiftRoutes := apiAdminRoutes.Group("/gifts")
			{
				apiGiftRoutes.Use(middlewares.AuthorizeJwt(db))
				apiGiftRoutes.Use(middlewares.AuthorizeAdminRole())
				routes.GiftRoutes(apiGiftRoutes, db)

			}

			apiVoucherRoutes := apiAdminRoutes.Group("/vouchers")
			{
				apiVoucherRoutes.Use(middlewares.AuthorizeJwt(db))
				apiVoucherRoutes.Use(middlewares.AuthorizeAdminRole())
				routes.VoucherRoutes(apiVoucherRoutes, db)
			}

			apiAccountRoutes := apiAdminRoutes.Group("/accounts")
			{
				apiAccountRoutes.Use(middlewares.AuthorizeJwt(db))
				apiAccountRoutes.Use(middlewares.AuthorizeAdminRole())
				routes.AccountRoutes(apiAccountRoutes, db)
			}

			apiTransaction := apiAdminRoutes.Group("/transactions")
			{
				apiTransaction.Use(middlewares.AuthorizeJwt(db))
				apiTransaction.Use(middlewares.AuthorizeAdminRole())
				routes.TransactionRoutes(apiTransaction, db)
			}
			apiExchangeRoutes := apiAdminRoutes.Group("/test_voucher")
			{
				apiExchangeRoutes.Use(middlewares.AuthorizeJwt(db))
				apiExchangeRoutes.Use(middlewares.AuthorizeAdminRole())
				routes.TestVoucher(apiExchangeRoutes, db)
			}
			apiTransferRoutes := apiAdminRoutes.Group("/transfer")
			{
				apiTransferRoutes.Use(middlewares.AuthorizeJwt(db))
				apiTransferRoutes.Use(middlewares.AuthorizeAdminRole())
				routes.TransferRoutes(apiTransferRoutes, db)
			}

		}

		apiAuthRoutes := apiRoutesV1.Group("/auth")
		{
			routes.AuthRoute(apiAuthRoutes, db)
		}
		apiExchangeRoutes := apiRoutesV1.Group("/exchange_voucher")
		{
			apiExchangeRoutes.Use(middlewares.AuthorizeJwt(db))
			routes.ExchangeVoucherRoutes(apiExchangeRoutes, db)
		}
	}

	server.Run(":" + port)
}
