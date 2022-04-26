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

	//db := connection.GetDB()
	// db.Create(&receipts.ReceiptItem{
	// 	ProductCode: "abcd",
	// 	Name:        "abcd",
	// 	UnitPrice:   10000,
	// 	Quantity:    1,
	// })
	// db.Create(&receipts.Receipt{
	// 	TransactionID:   "abcd",
	// 	TransactionDate: "abcd",
	// 	StatusID:        1,
	// })

	server.Run(":8080")
}
