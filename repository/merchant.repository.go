package repository

import (
	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantRepository interface{
	CreateMerchant(merchant entity.MerchantEntity) (entity.MerchantEntity, dto.ErrResponse)
	GetMerchantAll(user_id uuid.UUID) ([]entity.MerchantEntity, dto.ErrResponse)
	DetailMerchant(id uuid.UUID) (entity.MerchantEntity, dto.ErrResponse)
}

type merchantRepository struct {
	db *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepository{
		db: db,
	}
}

func (m *merchantRepository) CreateMerchant(merchant entity.MerchantEntity) (entity.MerchantEntity, dto.ErrResponse) {
	err := m.db.Create(&merchant).Error
	if err != nil {
		return merchant, dto.ErrResponse{Message: "Failed to Create Merchant"}
	}

	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "User with Number ID " + merchant.UserID.String() + " success created merchant with ID " + merchant.ID.String(),
	}).Error

	if err != nil {
		return merchant, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return merchant, dto.ErrResponse{Message: ""}
}

func (m *merchantRepository) GetMerchantAll(user_id uuid.UUID) ([]entity.MerchantEntity, dto.ErrResponse) {
	var merchants []entity.MerchantEntity
	err := m.db.Where("user_id = ?", user_id).Find(&merchants).Error

	if err != nil {
		return merchants, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "User with number ID " + user_id.String() + " success to get all merchant",
	}).Error

	if err != nil {
		return merchants, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return merchants, dto.ErrResponse{Message: ""}
}

func (m *merchantRepository) DetailMerchant(id uuid.UUID) (entity.MerchantEntity, dto.ErrResponse) {
	var merchant entity.MerchantEntity
	err := m.db.Joins("user_entities").Joins("rekening_entities").Where("id = ?", id).First(&merchant).Error

	if err != nil {
		return merchant, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Merchant with Number ID " + merchant.ID.String() + " get access",
	}).Error

	if err != nil {
		return merchant, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return merchant, dto.ErrResponse{Message: ""}
}

