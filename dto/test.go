package dto

type TestDTO struct {
	Secret string `form:"secret" binding:"required,gte=5"`
}
