package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/utils"
)

type Database interface {
	ConnectDB()
	Init()
	GetDB() *gorm.DB
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewDBConnection() Database {
	return &database{}
}

func (db *database) ConnectDB() {
	password := utils.GetDotEnv("DB_PASSWORD")

	dsn := "root:" + password + "@tcp(127.0.0.1:3306)/rpa_voucher_exchange?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.connection = d
}

func (db *database) Init() {
	db.connection.AutoMigrate(
		&entities.ReceiptStatus{},
		&entities.Receipt{},
		&entities.ReceiptItem{},
		&entities.Customer{},
		&entities.Voucher{},
		&entities.Product{},
		&entities.Gift{},
	)
}

func (db *database) GetDB() *gorm.DB {
	return db.connection
}

func (db *database) CloseDB() {
	sqlDB, err := db.connection.DB()
	if err != nil {
		panic("Failed to close connection")
	}
	sqlDB.Close()
}
