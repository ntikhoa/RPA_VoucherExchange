package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type GiftRepo interface {
	FindByID(giftID uint) (entities.Gift, error)
}

type giftRepo struct {
	db *gorm.DB
}

func NewGiftRepo(db *gorm.DB) GiftRepo {
	return &giftRepo{db: db}
}

func (r *giftRepo) FindByID(giftID uint) (entities.Gift, error) {
	gift := entities.Gift{
		Model: gorm.Model{
			ID: giftID,
		},
	}
	err := r.db.First(&gift).Error

	return gift, err
}
