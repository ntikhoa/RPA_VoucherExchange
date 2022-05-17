package dto

type LoginDTO struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
