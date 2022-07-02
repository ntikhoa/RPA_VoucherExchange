package dto

type RegisterDTO struct {
	RegisterSaleDTO
	ProviderID uint `form:"provider_id" json:"provider_id" binding:"required"`
}

func NewRegisterDTO(dto RegisterSaleDTO, providerID uint) RegisterDTO {
	return RegisterDTO{
		RegisterSaleDTO: dto,
		ProviderID:      providerID,
	}
}
