package dto

type RegisterDTO struct {
	Username          string `form:"username" binding:"required"`
	Password          string `form:"password" binding:"required,gte=6"`
	ConfirmedPassword string `form:"confirmed_password" binding:"required,gte=6"`
	Name              string `form:"name"`
	ProviderID        uint   `form:"provider_id" binding:"required"`
}
