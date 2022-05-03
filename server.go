package main

import (
	"github.com/RPA_VoucherExchange/configs"
	"github.com/gin-gonic/gin"
)

var (
	conn = configs.NewDBConnection()
)

func main() {
	server := gin.New()

	conn.ConnectDB()
	defer conn.CloseDB()
	conn.Init()

	server.Run(":8080")
}
