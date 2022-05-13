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
	port := os.Getenv("PORT")
	server := gin.New()

	conn.ConnectDB()
	defer conn.CloseDB()
	conn.Init()
	db := conn.GetDB()
	// utils.Seeding(db)

	server.Use(middlewares.SetContentTypeJSON())
	server.Use(middlewares.ErrorHandler())

	apiRoutesV1 := server.Group("/api/v1")
	{
		apiProductRoutes := apiRoutesV1.Group("/products")
		{
			apiProductRoutes.Use(middlewares.AuthorizeJwt())
			routes.ProductRoutes(apiProductRoutes, db)
		}
	}

	server.Run(":" + port)
}
