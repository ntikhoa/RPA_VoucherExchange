package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	"gorm.io/gorm"
)

type GiftService interface {
	FindByID(providerID uint, giftID uint) (entities.Gift, error)
}

type giftService struct {
	repo repositories.GiftRepo
}

func NewGiftService(repo repositories.GiftRepo) GiftService {
	return &giftService{repo: repo}
}

func (s *giftService) FindByID(providerID uint, giftID uint) (entities.Gift, error) {
	gift, err := s.repo.FindByID(giftID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gift, custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return gift, err
	}

	if gift.ProviderID != providerID {
		return gift, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return gift, nil
}
