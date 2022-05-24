package services

import (
	"errors"
	"math"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type GiftService interface {
	CreateGift(gift entities.Gift) error
	UpdateGift(gift entities.Gift) error
	DeleteGiftByID(giftID uint, providerID uint) error
	FindAllGiftWithPage(giftID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Gift, error)
	FindGiftByID(giftID uint, providerID uint) (entities.Gift, error)
	GetGiftCount(giftID uint) (int64, error)
}

type giftService struct {
	repo repositories.GiftRepo
}

func NewGiftService(repo repositories.GiftRepo) GiftService {
	return &giftService{
		repo: repo,
	}
}

func (s *giftService) CreateGift(gift entities.Gift) error {
	return s.repo.CreateGift(gift)
}

func (s *giftService) UpdateGift(gift entities.Gift) error {
	fetchedGift, err := s.repo.FindGiftByID(gift.Model.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}
	if fetchedGift.ProviderID != gift.ProviderID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return s.repo.UpdateGift((gift))
}

func (s *giftService) DeleteGiftByID(giftID uint, providerID uint) error {
	gift, err := s.repo.FindGiftByID(giftID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if gift.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return s.repo.DeleteGiftByID(giftID)
}

func (s *giftService) FindAllGiftWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Gift, error) {
	pagingMetadata := viewmodel.PagingMetadata{}

	count, err := s.repo.GetGiftCount(providerID)
	if err != nil {
		return pagingMetadata, nil, err
	}
	d := float64(count) / float64(perPage)
	totalPages := int(math.Ceil(d))
	if page > totalPages {
		return pagingMetadata, nil, custom_error.NewNotFoundError(constants.EXHAUSTED_ERROR)
	}

	pagingMetadata = viewmodel.PagingMetadata{
		Page:         page,
		PerPage:      perPage,
		TotalPages:   totalPages,
		TotalRecords: int(count),
	}
	gifts, err := s.repo.FindAllGiftWithPage(providerID, page, perPage)
	return pagingMetadata, gifts, err
}

func (s *giftService) FindGiftByID(giftID uint, providerID uint) (entities.Gift, error) {
	gift, err := s.repo.FindGiftByID(giftID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gift, &custom_error.NotFoundError{}
		}
		return gift, err
	}

	if gift.ProviderID != providerID {
		return gift, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return gift, nil
}

func (s *giftService) GetGiftCount(providerID uint) (int64, error) {
	return s.repo.GetGiftCount(providerID)
}
