package repositories

import (
	"github.com/RPA_VoucherExchange/entities"
	"gorm.io/gorm"
)

type GiftRepo interface {
	Create(gift entities.Gift) error
	Update(gift entities.Gift) error
	DeleteByID(giftID uint) error
	DeleteByIDs(giftIDs []uint) error
	FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Gift, error)
	FindByID(giftID uint) (entities.Gift, error)
	FindByIDs(giftIDs []uint) ([]entities.Gift, error)
	Search(query string, providerID uint) ([]entities.Gift, error)
	Count(providerID uint) (int64, error)
	GetAll(providerID uint) ([]entities.Gift, error)
	CheckExistence(providerID uint, giftIDs []uint) ([]uint, error)
}

type giftRepo struct {
	db *gorm.DB
}

func NewGiftRepo(db *gorm.DB) GiftRepo {
	return &giftRepo{
		db: db,
	}
}

func (repo *giftRepo) Create(gift entities.Gift) error {
	return repo.db.Create(&gift).Error
}

func (repo *giftRepo) Update(gift entities.Gift) error {
	return repo.db.Omit("created_at").Save(&gift).Error
}

func (repo *giftRepo) DeleteByID(giftID uint) error {
	return repo.db.
		Delete(&entities.Gift{}, giftID).
		Error
}

func (repo *giftRepo) FindAllWithPage(providerID uint, page int, perPage int) ([]entities.Gift, error) {
	gifts := []entities.Gift{}
	tx := repo.db.
		Where(&entities.Gift{ProviderID: providerID}).
		Limit(perPage).
		Offset((page - 1) * perPage).
		Find(&gifts)
	return gifts, tx.Error
}

func (repo *giftRepo) FindByID(giftID uint) (entities.Gift, error) {
	gift := entities.Gift{}
	err := repo.db.First(&gift, giftID).Error

	return gift, err
}

func (repo *giftRepo) Search(query string, providerID uint) ([]entities.Gift, error) {
	gifts := []entities.Gift{}
	err := repo.db.
		Where("provider_id = ? AND gift_name LIKE ?", providerID, "%"+query+"%").
		Find(&gifts).
		Error
	return gifts, err
}

func (repo *giftRepo) Count(providerID uint) (int64, error) {
	var count int64
	err := repo.db.
		Model(&entities.Gift{ProviderID: providerID}).
		Count(&count).
		Error

	return count, err
}

func (repo *giftRepo) GetAll(providerID uint) ([]entities.Gift, error) {
	var gifts []entities.Gift

	err := repo.
		db.
		Where(&entities.Gift{ProviderID: providerID}).
		Order("gift_name").
		Find(&gifts).
		Error

	return gifts, err
}

func (repo *giftRepo) FindByIDs(giftIDs []uint) ([]entities.Gift, error) {
	var gifts []entities.Gift
	err := repo.db.Find(&gifts, giftIDs).Error
	return gifts, err
}

func (repo *giftRepo) DeleteByIDs(giftIDs []uint) error {
	var gifts []entities.Gift
	return repo.db.
		Delete(&gifts, giftIDs).
		Error
}

type giftID struct {
	ID uint
}

func (repo *giftRepo) CheckExistence(providerID uint, giftIDs []uint) ([]uint, error) {
	var IDs []giftID
	tx := repo.
		db.
		Model(&entities.Gift{}).
		Where("id IN (?) AND provider_id = ?", giftIDs, providerID).
		Find(&IDs)

	ids := []uint{}

	for _, id := range IDs {
		ids = append(ids, id.ID)
	}
	return ids, tx.Error
}
