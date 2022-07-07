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

func SeedProducts(db *gorm.DB, variant int) {
	providerID := []uint{1, 4}
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
		products = append(products, entities.Product{ProductName: productName[i], ProviderID: providerID[variant]})
	}
	db.CreateInBatches(&products, len(products))
}

func SeedVouchers(db *gorm.DB, variant int) {
	providerID := []uint{1, 4}

	productsJunkFood := map[string][]uint{
		"XUC XICH DUC SMG 500": {84, 774},
		"XUC XICH KABANA LC 5": {85, 784},
		"PHO MAI BEGA TASTY":   {86, 794},
	}

	productsSoftDrink := map[string][]uint{
		"MRD SAXI 320/330ML": {80, 744},
		"COCA ZERO 320ML":    {81, 754},
		"PEPSI COLA 1.5L":    {83, 764},
	}

	productTampon := map[string][]uint{
		"BVS DIANA STCF8M":    {106, 834},
		"BVS DIANA HANG NGAY": {107, 844},
	}

	vouchers := []entities.Voucher{
		{
			Name: "Junk Food 2nd",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Xuc Xich Duc SMG 500, Xuc Xich Kabana LC 5, Pho Mai Bega Tasty\n
			Từ 100,000 vnd trở lên: tặng một Bánh Mì Sandwich
			`,
			TotalPriceMin: 100000,
			TotalPriceMax: 200000,
			Total:         100,
			Remaining:     100,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productsJunkFood, variant),
			Gift:          entities.Gift{GiftName: "Bánh Mì Sandwich", ProviderID: providerID[variant]},
		},
		{
			Name: "Junk Food 1st",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Xuc Xich Duc SMG 500, Xuc Xich Kabana LC 5, Pho Mai Bega Tasty\n
			Từ 200,000 vnd trở lên: tặng một bộ dao + nĩa trị giá 100,000vnđ
			`,
			TotalPriceMin: 200000,
			TotalPriceMax: 0,
			Total:         50,
			Remaining:     50,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productsJunkFood, variant),
			Gift:          entities.Gift{GiftName: "Một bộ dao + nĩa trị giá 100,000vnđ", ProviderID: providerID[variant]},
		},
		{
			Name: "Soft Drink",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Coca Zero, Mirinda Saxi, Pepsi Cola\n
			Từ 20,000 vnd trở lên: tặng một bộ ống hút thủy tinh
			`,
			TotalPriceMin: 20000,
			TotalPriceMax: 0,
			Total:         100,
			Remaining:     100,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productsSoftDrink, variant),
			Gift:          entities.Gift{GiftName: "Một bộ ống hút thủy tinh", ProviderID: providerID[variant]},
		},
		{
			Name: "Tampon 2nd",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Diana STCF8M hoặc Diana hằng ngày\n
			Từ 49,000 vnd trở lên: tặng một khăn ướt chăm sóc da Fressi Care Cool
			`,
			TotalPriceMin: 49000,
			TotalPriceMax: 79000,
			Total:         100,
			Remaining:     100,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productTampon, variant),
			Gift:          entities.Gift{GiftName: "Một khăn ướt chăm sóc da Fressi Care Cool", ProviderID: providerID[variant]},
		},
		{
			Name: "Tampon 1st",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Diana STCF8M hoặc Diana hằng ngày\n
			Từ 79,000 vnd trở lên: tặng một máy hút bụi cầm tay
			`,
			TotalPriceMin: 79000,
			TotalPriceMax: 0,
			Total:         50,
			Remaining:     50,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productTampon, variant),
			Gift:          entities.Gift{GiftName: "Một máy hút bụi cầm tay", ProviderID: providerID[variant]},
		},
	}

	for _, voucher := range vouchers {
		db.Create(&voucher)
	}
	// db.CreateInBatches(&vouchers, len(vouchers))
}

func createProductsForVoucher(productsName map[string][]uint, variant int) []entities.Product {
	var products []entities.Product
	for _, value := range productsName {
		products = append(products, entities.Product{
			Model: gorm.Model{
				ID: value[variant],
			},
		})
	}
	return products
}
