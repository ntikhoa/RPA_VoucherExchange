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

func SeedGift(db *gorm.DB, variant int) {
	providerID := []uint{1, 4}
	giftName := []string{
		"Bánh Mì Sandwich",
		"Một bộ dao + nĩa trị giá 100,000vnđ",
		"Một bộ ống hút thủy tinh",
		"Một máy hút bụi cầm tay",
		"Một khăn ướt chăm sóc da Fressi Care Cool",
	}

	gifts := []entities.Gift{}

	for i := 0; i < len(giftName); i++ {
		gifts = append(gifts, entities.Gift{GiftName: giftName[i], ProviderID: providerID[variant]})
	}
	db.CreateInBatches(&gifts, len(gifts))
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

	giftJunkFood2nd := []uint{53, 264}
	giftJunkFood1st := []uint{54, 274}
	SoftDrink := []uint{55, 284}
	Tampon2nd := []uint{56, 294}
	Tampon1st := []uint{57, 304}

	vouchers := []entities.Voucher{
		{
			Name: "Junk Food 2nd",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Xuc Xich Duc SMG 500, Xuc Xich Kabana LC 5, Pho Mai Bega Tasty\n
			Từ 100,000 vnd trở lên: tặng một Bánh Mì Sandwich
			`,
			TotalPriceMin: 100000,
			TotalPriceMax: 200000,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productsJunkFood, variant),
			GiftID:        giftJunkFood2nd[variant],
		},
		{
			Name: "Junk Food 1st",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Xuc Xich Duc SMG 500, Xuc Xich Kabana LC 5, Pho Mai Bega Tasty\n
			Từ 200,000 vnd trở lên: tặng một bộ dao + nĩa trị giá 100,000vnđ
			`,
			TotalPriceMin: 200000,
			TotalPriceMax: 0,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productsJunkFood, variant),
			GiftID:        giftJunkFood1st[variant],
		},
		{
			Name: "Soft Drink",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Coca Zero, Mirinda Saxi, Pepsi Cola\n
			Từ 20,000 vnd trở lên: tặng một bộ ống hút thủy tinh
			`,
			TotalPriceMin: 20000,
			TotalPriceMax: 0,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productsSoftDrink, variant),
			GiftID:        SoftDrink[variant],
		},
		{
			Name: "Tampon 2nd",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Diana STCF8M hoặc Diana hằng ngày\n
			Từ 49,000 vnd trở lên: tặng một khăn ướt chăm sóc da Fressi Care Cool
			`,
			TotalPriceMin: 49000,
			TotalPriceMax: 79000,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productTampon, variant),
			GiftID:        Tampon2nd[variant],
		},
		{
			Name: "Tampon 1st",
			Description: `Hóa đơn có một hoặc nhiều mặt hàng trong danh sách sau:\n
			Diana STCF8M hoặc Diana hằng ngày\n
			Từ 79,000 vnd trở lên: tặng một máy hút bụi cầm tay
			`,
			TotalPriceMin: 79000,
			TotalPriceMax: 0,
			Published:     true,
			ProviderID:    providerID[variant],
			Products:      createProductsForVoucher(productTampon, variant),
			GiftID:        Tampon1st[variant],
		},
	}

	for _, voucher := range vouchers {
		db.Create(&voucher)
	}
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
