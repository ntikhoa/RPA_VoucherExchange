package entities

import "gorm.io/gorm"

type ReceiptImage struct {
	gorm.Model
	Url       string `gorm:"type:varchar(255); not null"`
	ReceiptID uint
}

func NewReceiptImage(fileName string) ReceiptImage {
	return ReceiptImage{
		Url: fileName,
	}
}

func NewReceiptImages(filesNames []string) []ReceiptImage {
	var receiptImages []ReceiptImage
	for _, fileName := range filesNames {
		receiptImages = append(receiptImages, NewReceiptImage(fileName))
	}
	return receiptImages
}
