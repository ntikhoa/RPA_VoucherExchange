package main

import (
	"os"
	"time"

	"github.com/RPA_VoucherExchange/configs"
	"github.com/RPA_VoucherExchange/middlewares"
	"github.com/RPA_VoucherExchange/routes"
	"github.com/gin-contrib/cors"
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
	// conn.Init()
	db := conn.GetDB()
	// utils.Seeding(db)

	server.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"},
		ExposeHeaders: []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Max"},
		MaxAge:        12 * time.Hour,
	}))

	// server.Use(middlewares.SetHeader())
	server.Use(middlewares.ErrorHandler())

	apiRoutesV1 := server.Group("/api/v1")
	{
		apiProductRoutes := apiRoutesV1.Group("/products")
		{
			apiProductRoutes.Use(middlewares.AuthorizeJwt())
			routes.ProductRoutes(apiProductRoutes, db)
		}

		apiAuthRoutes := apiRoutesV1.Group("/auth")
		{
			routes.AuthRoutes(apiAuthRoutes, db)
		}

		apiRoutesV1.GET("/test", middlewares.AuthorizeJwt())
	}

	server.Run(":" + port)
}
