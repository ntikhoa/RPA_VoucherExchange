package main

import (
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
	//for loading local .env file
	// utils.LoadDotEnv()

	port := os.Getenv("PORT")
	server := gin.New()

	conn.ConnectDB()
	defer conn.CloseDB()
	conn.Init()
	db := conn.GetDB()
	// utils.SeedProducts(db)

	server.Use(middlewares.SetHeader())
	server.Use(middlewares.ErrorHandler())

	apiRoutesV1 := server.Group("/api/v1")
	{
		apiProductRoutes := apiRoutesV1.Group("/products")
		{
			apiProductRoutes.Use(middlewares.AuthorizeJwt(db))
			routes.ProductRoutes(apiProductRoutes, db)
		}

		apiAuthRoutes := apiRoutesV1.Group("/auth")
		{
			routes.AuthRoutes(apiAuthRoutes, db)
		}

		apiVoucherRoutes := apiRoutesV1.Group("/vouchers")
		{
			apiVoucherRoutes.Use(middlewares.AuthorizeJwt(db))
			routes.VoucherRoutes(apiVoucherRoutes, db)
		}
	}

	server.Run(":" + port)
}
