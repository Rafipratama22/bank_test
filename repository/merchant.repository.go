package repository

import (
	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	CreateMerchant(merchant entity.MerchantEntity) (entity.MerchantEntity, dto.ErrResponse)
	GetMerchantAll(user_id uuid.UUID) ([]dto.MerchantListResponse, dto.ErrResponse)
	DetailMerchant(id uuid.UUID) (dto.MerchantListResponse, dto.ErrResponse)
	UpdateMerchant(id uuid.UUID, merchant entity.MerchantEntity) (entity.MerchantEntity, dto.ErrResponse)
	DeleteMerchant(id uuid.UUID, userId uuid.UUID) (dto.ErrResponse)
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

func (m *merchantRepository) GetMerchantAll(user_id uuid.UUID) ([]dto.MerchantListResponse, dto.ErrResponse) {
	var merchants []entity.MerchantEntity
	var user entity.UserEntity
	var rekening entity.RekeningEntity
	var response []dto.MerchantListResponse
	err := m.db.Where("user_id = ?", user_id).Find(&merchants).Error
	if err != nil {
		return response, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	for i := 0; i < len(merchants); i++ {
		m.db.Model(&user).Where("id = ?", merchants[i].UserID).First(&user)
		m.db.Model(&rekening).Where("no_rekening = ?", merchants[i].NoRekening).First(&rekening)
		responseOne := dto.MerchantListResponse{
			ID:         merchants[i].ID,
			Name:       merchants[i].Name,
			Address:    merchants[i].Address,
			Phone:      merchants[i].Phone,
			NoRekening: merchants[i].NoRekening,
			UserID:     merchants[i].UserID,
			User:       user,
			Rekening:   rekening,
		}
		response = append(response, responseOne)
	}

	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "User with number ID " + user_id.String() + " success to get all merchant",
	}).Error

	if err != nil {
		return response, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return response, dto.ErrResponse{Message: ""}
}

func (m *merchantRepository) DetailMerchant(id uuid.UUID) (dto.MerchantListResponse, dto.ErrResponse) {
	var merchant entity.MerchantEntity
	var user entity.UserEntity
	var rekening entity.RekeningEntity
	var response dto.MerchantListResponse
	err := m.db.Model(&merchant).Where("id = ?", id).First(&merchant).Error

	if err != nil {
		return response, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	err = m.db.Model(&user).Where("id = ?", merchant.UserID).First(&user).Error

	if err != nil {
		return response, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	err = m.db.Model(&rekening).Where("no_rekening = ?", merchant.NoRekening).First(&rekening).Error

	if err != nil {
		return response, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	response.User = user
	response.Rekening = rekening
	response.ID = merchant.ID
	response.Name = merchant.Name
	response.Address = merchant.Address
	response.Phone = merchant.Phone
	response.NoRekening = merchant.NoRekening
	response.UserID = merchant.UserID
	
	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Merchant with Number ID " + merchant.ID.String() + " get access",
	}).Error

	if err != nil {
		return response, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return response, dto.ErrResponse{Message: ""}
}

func (m *merchantRepository) UpdateMerchant(id uuid.UUID, merchant entity.MerchantEntity) (entity.MerchantEntity, dto.ErrResponse) {
	var merchantOld entity.MerchantEntity
	err := m.db.Model(&merchantOld).Where("id = ?", id).Find(&merchantOld).Error

	if err != nil {
		panic(err)
		// return merchant, dto.ErrResponse{Message: "Failed to Get Merchant"}
	}

	if merchantOld.UserID != merchant.UserID {
		return merchant, dto.ErrResponse{Message: "Not Allowed"}
	}
	err = m.db.Model(&merchant).Where("id = ?", id).Updates(merchant).Error

	if err != nil {
		return merchant, dto.ErrResponse{Message: "Failed to Update Merchant"}
	}
	
	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Merchant with Number ID " + id.String() + " success update",
	}).Error

	if err != nil {
		return merchant, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return merchant, dto.ErrResponse{Message: ""}
}

func (m *merchantRepository) DeleteMerchant(id uuid.UUID, userId uuid.UUID) (dto.ErrResponse) {
	var merchant entity.MerchantEntity
	var merchantOld entity.MerchantEntity
	err := m.db.Model(&merchant).Where("id = ?", id).Find(&merchantOld).Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Get Merchant"}
	}
	if merchantOld.UserID != userId {
		return dto.ErrResponse{Message: "Not Allowed"}
	}

	err = m.db.Model(&merchant).Where("id = ?", id).Delete(&merchant).Error

	if err != nil {
		return dto.ErrResponse{Message: "Failed to Delete Merchant"}
	}

	err = m.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Merchant with Number ID " + id.String() + " success delete",
	}).Error

	if err != nil {
		return dto.ErrResponse{Message: "Failed to Create History"}
	}

	return dto.ErrResponse{Message: ""}
}