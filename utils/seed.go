package utils

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) {
	db.Create(&entities.Provider{
		Name: "CONG TY TNHH CONG DONG VIET",
	})
}

func SeedProducts(db *gorm.DB) {
	productName := []string{
		"TN-MAM LUA MACH",
		"GUM MFA ECA VI",
		"SNACK OS TOM 39G",
		"SCU YAKULT (65MLX10)",
		"L6 STINGRED SLK320ML",
		"MRD SAXI 320/330ML",
		"COCA ZERO 320ML",
		"PEPSI COLA 1.5L",
		"XUC XICH DUC SMG 500",
		"XUC XICH KABANA LC 5",
		"PHO MAI BEGA TASTY",
		"HAU NUONG HASUBI200G",
		"KEM 3IN1 WALL 390g",
		"AJI-MAYO 130G",
		"BVS DIANA STCF8M",
		"BVS DIANA HANG NGAY",
		"P/S NSM ACTIVE DEFEN",
		"KDR CLOSE UP 230G",
		"BVS HN DIANA",
		"KHAY DA TS-3653",
	}

	products := []entities.Product{}

	for i := 0; i < len(productName); i++ {
		products = append(products, entities.Product{ProductName: productName[i], ProviderID: 4})
	}
	db.CreateInBatches(&products, len(products))
}
