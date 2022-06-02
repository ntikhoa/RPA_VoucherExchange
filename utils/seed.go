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
		"TV NUOC UONG 1.5L",
		"TN-MAM LUA MACH",
		"NUOC UONG 500ML",
		"GUM MFA ECA VI",
		"SNACK OS TOM 39G",
		"SCU YAKULT (65MLX10)",
		"PHAN TRUNG QUE",
		"L6 STINGRED SLK320ML",
		"MRD SAXI 320/330ML",
		"COCA ZERO 320ML",
		"THACHDUAAHONG90GX6",
		"PEPSI COLA 1.5L",
		"XUC XICH DUC SMG 500",
		"XUC XICH KABANA LC 5",
		"PHO MAI BEGA TASTY",
		"CA LOC FILLET VG",
		"SHASHIMI FILLET CA HOI",
		"LUON CA HOI BBQ",
		"COMBO LAU HAISAN RAU",
		"HAU NUONG HASUBI200G",
		"KEM 3IN1 WALL 390g",
		"AJI-MAYO 130G",
		"CA CHUA SOCOLA 250G",
		"THANH LONG RUOT DO",
		"RONG NHO TUOI 100G",
		"BVS DIANA STCF8M",
		"BVS DIANA HANG NGAY",
		"P/S NSM ACTIVE DEFEN",
		"KDR CLOSE UP 230G",
		"BVS HN DIANA",
		"THUOC DUONG HOA H.10",
		"NUOC EP OI CHAI 330M",
		"TV HOP KHU MUI140G",
		"KHAY DA TS-3653",
	}

	products := []entities.Product{}

	for i := 0; i < len(productName); i++ {
		products = append(products, entities.Product{ProductName: productName[i], ProviderID: 4})
	}
	db.CreateInBatches(&products, len(products))
}
