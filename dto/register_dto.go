package dto

type RegisterDTO struct {
	Username          string `form:"username" json:"username" binding:"required"`
	Password          string `form:"password" json:"password" binding:"required,gte=6"`
	ConfirmedPassword string `form:"confirmed_password" json:"confirmed_password" binding:"required,gte=6"`
	Name              string `form:"name" json:"name"`
	ProviderID        uint   `form:"provider_id" json:"provider_id" binding:"required"`
}
