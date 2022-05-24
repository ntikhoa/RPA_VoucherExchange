package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type GiftRepo interface {
	CreateGift(gift entities.Gift) error
	UpdateGift(gift entities.Gift) error
	DeleteGiftByID(giftID uint) error
	FindAllGiftWithPage(providerID uint, page int, perPage int) ([]entities.Gift, error)
	FindGiftByID(giftID uint) (entities.Gift, error)
	GetGiftCount(providerID uint) (int64, error)
}

type giftRepo struct {
	db *gorm.DB
}

func NewGiftRepo(db *gorm.DB) GiftRepo {
	return &giftRepo{
		db: db,
	}
}

func (repo *giftRepo) CreateGift(gift entities.Gift) error {
	return repo.db.Create(&gift).Error
}

func (repo *giftRepo) UpdateGift(gift entities.Gift) error {
	return repo.db.Save(&gift).Error
}

func (repo *giftRepo) DeleteGiftByID(giftID uint) error {
	return repo.db.Unscoped().Delete(&entities.Gift{}, giftID).Error
}

func (repo *giftRepo) FindAllGiftWithPage(providerID uint, page int, perPage int) ([]entities.Gift, error) {
	gifts := []entities.Gift{}
	tx := repo.db.
		Where(&entities.Gift{ProviderID: providerID}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&gifts)
	return gifts, tx.Error
}

func (repo *giftRepo) FindGiftByID(giftID uint) (entities.Gift, error) {
	gift := entities.Gift{
		Model: gorm.Model{
			ID: giftID,
		},
	}
	err := repo.db.First(&gift).Error

	return gift, err
}

func (repo *giftRepo) GetGiftCount(providerID uint) (int64, error) {
	var count int64
	err := repo.db.
		Model(&entities.Gift{ProviderID: providerID}).
		Count((&count)).
		Error

	return count, err
}
