package main

import (
	"os"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/routes"
	"github.com/RPA_VoucherExchange/utils"
	"github.com/gin-gonic/gin"
)

var (
	conn = configs.NewDBConnection()
)

func main() {
	//for loading local .env file
	utils.LoadDotEnv()

	port := os.Getenv("PORT")
	server := gin.New()

	conn.ConnectDB()
	defer conn.CloseDB()
	// conn.Init()
	db := conn.GetDB()
	// utils.SeedProducts(db)

	server.Use(middlewares.SetHeader())
	server.Use(middlewares.ErrorHandler())

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
		}

		apiAuthRoutes := apiRoutesV1.Group("/auth")
		{
			routes.AuthRoute(apiAuthRoutes, db)
		}
		apiExchangeRoutes := apiRoutesV1.Group("/exchange_voucher")
		{
			routes.ExchangeVoucherRoutes(apiExchangeRoutes)
		}
	}

	server.Run(":" + port)
}
