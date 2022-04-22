package main

import (
	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

var (
	connection = configs.NewDBConnection()
)

func main() {
	server := gin.New()

	connection.ConnectDB()
	defer connection.CloseDB()
	connection.Init()

	server.Run(":8080")
}
