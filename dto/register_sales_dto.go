package dto

type RegisterSaleDTO struct {
	Username          string `form:"username" json:"username" binding:"required"`
	Password          string `form:"password" json:"password" binding:"required,gte=6"`
	ConfirmedPassword string `form:"confirmed_password" json:"confirmed_password" binding:"required,gte=6"`
	Name              string `form:"name" json:"name"`
}

func (dto RegisterSaleDTO) ToRegisterDTO(providerID uint) RegisterDTO {
	return RegisterDTO{
		Username:          dto.Username,
		Password:          dto.Password,
		ConfirmedPassword: dto.ConfirmedPassword,
		ProviderID:        providerID,
		Name:              dto.Name,
	}
}
