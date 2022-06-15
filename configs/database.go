package configs

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/RPA_VoucherExchange/entities"
)

type Database interface {
	ConnectDB()
	Init()
	GetDB() *gorm.DB
	CloseDB()

	getURL() string
}

type database struct {
	connection *gorm.DB
}

func NewDBConnection() Database {
	return &database{}
}

func (db *database) ConnectDB() {
	dsn := db.getURL()
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.connection = d
}

func (db *database) Init() {
	db.connection.AutoMigrate(
		&entities.Provider{},
		&entities.ReceiptStatus{},
		&entities.Receipt{},
		&entities.ReceiptItem{},
		&entities.Customer{},
		&entities.Voucher{},
		&entities.Product{},
		&entities.Gift{},
		&entities.Employee{},
		&entities.Role{},
	)
	initEnum(db.connection)
}

func initEnum(db *gorm.DB) {
	db.Create(&entities.Role{
		Model: gorm.Model{
			ID: 1,
		},
		Description: "ADMIN",
	})

	db.Create(&entities.Role{
		Model: gorm.Model{
			ID: 2,
		},
		Description: "SALE",
	})

	db.Create(&entities.ReceiptStatus{
		Model: gorm.Model{
			ID: 1,
		},
		Description: "COMPLETE",
	})

	db.Create(&entities.ReceiptStatus{
		Model: gorm.Model{
			ID: 2,
		},
		Description: "PENDING",
	})
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

func (db *database) getURL() string {
	//for local db instance
	password := os.Getenv("LOCAL_DB_PASSWORD")
	dsn := "root:" + password + "@tcp(127.0.0.1:3306)/rpa_voucher_exchange?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn

	//for remote db instance
	// username := os.Getenv("REMOTE_DB_USERNAME")
	// password := os.Getenv("REMOTE_DB_PASSWORD")
	// hostname := "@tcp(" + os.Getenv("REMOTE_DB_HOST") + ")"
	// dbName := os.Getenv("REMOTE_DB_NAME")
	// option := "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := username + ":" + password + hostname + "/" + dbName + option
	// return dsn
}
