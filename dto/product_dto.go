package dto

type ProductDTO struct {
	ProductName string `form:"product_name" json:"product_name" binding:"required"`
}
