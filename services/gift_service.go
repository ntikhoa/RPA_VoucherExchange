package services

import (
	"errors"

	"github.com/RPA_VoucherExchange/constants"
	"github.com/RPA_VoucherExchange/custom_error"
	"github.com/RPA_VoucherExchange/dto"
	"github.com/RPA_VoucherExchange/entities"
	"github.com/RPA_VoucherExchange/repositories"
	viewmodel "github.com/RPA_VoucherExchange/view_model"
	"gorm.io/gorm"
)

type GiftService interface {
	Create(giftDTO dto.GiftDTO, providerID uint) error
	Update(giftDTO dto.GiftDTO, providerID uint, giftID uint) error
	DeleteByID(giftID uint, providerID uint) error
	DeleteByIDs(giftIDs []uint, providerID uint) error
	FindAllWithPage(
		providerID uint,
		page int,
		perPage int) (viewmodel.PagingMetadata, []entities.Gift, error)
	FindByID(giftID uint, providerID uint) (entities.Gift, error)
	Search(query string, providerID uint) ([]entities.Gift, error)
	GetCount(providerID uint) (int64, error)
	GetAll(providerID uint) ([]entities.Gift, error)
	CheckExistence(providerID uint, giftIDs []uint) error
}

type giftService struct {
	repo repositories.GiftRepo
}

func NewGiftService(repo repositories.GiftRepo) GiftService {
	return &giftService{
		repo: repo,
	}
}

func (s *giftService) Create(giftDTO dto.GiftDTO, providerID uint) error {
	gift := entities.NewGift(giftDTO, providerID)
	return s.repo.Create(gift)
}

func (s *giftService) Update(giftDTO dto.GiftDTO, providerID uint, giftID uint) error {
	gift := entities.NewGift(giftDTO, providerID)
	gift.Model.ID = giftID
	fetchedGift, err := s.repo.FindByID(gift.Model.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return custom_error.NewNotFoundError(constants.NOT_FOUND_ERROR)
		}
		return err
	}

	if fetchedGift.ProviderID != gift.ProviderID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return s.repo.Update(gift)
}

func (s *giftService) DeleteByID(giftID uint, providerID uint) error {
	gift, err := s.repo.FindByID(giftID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	if gift.ProviderID != providerID {
		return custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}
	return s.repo.DeleteByID(giftID)
}

func (s *giftService) FindAllWithPage(providerID uint, page int, perPage int) (viewmodel.PagingMetadata, []entities.Gift, error) {
	pagingMetadata := viewmodel.PagingMetadata{}

	pagingMetadata, err := paging(s.repo.Count, providerID, page, perPage)
	if err != nil {
		return pagingMetadata, nil, err
	}

	gifts, err := s.repo.FindAllWithPage(providerID, page, perPage)
	return pagingMetadata, gifts, err
}

func (s *giftService) FindByID(giftID uint, providerID uint) (entities.Gift, error) {
	gift, err := s.repo.FindByID(giftID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gift, custom_error.NewNotFoundError("gift " + constants.NOT_FOUND_ERROR)
		}
		return gift, err
	}

	if gift.ProviderID != providerID {
		return gift, custom_error.NewForbiddenError(constants.AUTHORIZE_ERROR)
	}

	return gift, nil
}

func (s *giftService) Search(query string, providerID uint) ([]entities.Gift, error) {
	gifts, err := s.repo.Search(query, providerID)
	if err != nil {
		return gifts, err
	}
	return gifts, nil
}

func (s *giftService) GetCount(providerID uint) (int64, error) {
	return s.repo.Count(providerID)
}

func (s *giftService) GetAll(providerID uint) ([]entities.Gift, error) {
	gifts, err := s.repo.GetAll(providerID)
	if err != nil {
		return nil, err
	}

	return gifts, nil
}

func (s *giftService) DeleteByIDs(giftIDs []uint, providerID uint) error {
	gifts, err := s.repo.FindByIDs(giftIDs)
	if err != nil {
		return err
	}
	var filteredGiftIDs []uint
	if len(gifts) > 0 {
		for _, gift := range gifts {
			if gift.ProviderID == providerID {
				filteredGiftIDs = append(filteredGiftIDs, gift.Model.ID)
			}
		}
	}
	if len(filteredGiftIDs) > 0 {
		return s.repo.DeleteByIDs(filteredGiftIDs)
	}
	return nil
}

func (s *giftService) CheckExistence(providerID uint, giftIDs []uint) error {
	fetchedID, err := s.repo.CheckExistence(providerID, giftIDs)
	if err != nil {
		return err
	}

	invalidProductIDs := extractInvalidIDs(fetchedID, giftIDs)

	if len(invalidProductIDs) > 0 {
		invalidIDstr := convertToStringError(invalidProductIDs)
		return custom_error.NewConflictError("invalid gift ids: " + invalidIDstr)
	}
	return nil
}
