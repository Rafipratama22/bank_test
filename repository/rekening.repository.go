package repository

import (
	"strconv"

	"github.com/Rafipratama22/bank_test/dto"
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type RekeningRepository interface {
	CreateRekening(saving entity.RekeningEntity) dto.ErrResponse
	UpdateRekening(id int, balance int, pin string) dto.ErrResponse
	GetRekeningAll(user_id uuid.UUID) ([]entity.RekeningEntity, dto.ErrResponse)
	GetRekeningByID(id int) (entity.RekeningEntity, dto.ErrResponse)
	TransferRekening(no_rekening uuid.UUID, balance int, pin string, id int) dto.ErrResponse
}

type rekeningRepository struct {
	db *gorm.DB
}

func NewRekeningRepository(db *gorm.DB) RekeningRepository {
	return &rekeningRepository{
		db: db,
	}
}

func (s *rekeningRepository) CreateRekening(saving entity.RekeningEntity) dto.ErrResponse {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saving.Pin), bcrypt.DefaultCost)
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Generate Password"}
	}
	saving.Pin = string(hashedPassword)
	err = s.db.Create(&saving).Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Create Rekening"}
	}
	err = s.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Create Rekening from " + saving.NoRekening.String(),
	}).Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Create History"}
	}
	return dto.ErrResponse{Message: ""}
}

func (s *rekeningRepository) UpdateRekening(id int, balance int, pin string) dto.ErrResponse {
	var rekening entity.RekeningEntity
	err := s.db.Where("id = ?", id).First(&rekening).Error
	if err != nil {
		return dto.ErrResponse{Message: "There is no rekening with this id"}
	}
	if rekening.Chance > 3 {
		return dto.ErrResponse{Message: "You have reached the maximum chance"}
	}
	err = bcrypt.CompareHashAndPassword([]byte(rekening.Pin), []byte(pin))
	if err != nil {
		err = s.db.Model(&rekening).Update("chance", rekening.Chance+1).Error
		if err != nil {
			return dto.ErrResponse{Message: "Failed to Update Rekening"}
		}
		return dto.ErrResponse{Message: "Wrong Pin"}
	}
	balanceNow := rekening.Balance + balance
	err = s.db.Model(&rekening).Updates(entity.RekeningEntity{
		Balance: balanceNow,
		Chance:  0,
	}).Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Update Rekening"}
	}

	err = s.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Rekening from " + rekening.NoRekening.String() + " has been updated",
	}).Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed to Create History"}
	}

	return dto.ErrResponse{Message: ""}
}

func (s *rekeningRepository) GetRekeningAll(user_id uuid.UUID) ([]entity.RekeningEntity, dto.ErrResponse) {
	var rekenings []entity.RekeningEntity
	var errResponse dto.ErrResponse
	err := s.db.Where("user_id = ?", user_id).Find(&rekenings).Error
	if err != nil {
		errResponse.Message = "There is no rekening with this user_id"
		return rekenings, errResponse
	}

	err = s.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "All Rekening from user " + user_id.String() + " has been accessed",
	}).Error
	
	if err != nil {
		return rekenings, dto.ErrResponse{Message: "Failed to Create History"}
	}

	errResponse.Message = ""
	return rekenings, errResponse
}

func (s *rekeningRepository) GetRekeningByID(id int) (entity.RekeningEntity, dto.ErrResponse) {
	var rekening entity.RekeningEntity
	err := s.db.Where("id = ?", id).First(&rekening).Error
	if err != nil {
		return rekening, dto.ErrResponse{Message: "There is no rekening with this id"}
	}
	err = s.db.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Detail Rekening from " + rekening.NoRekening.String() + " has been accessed",
	}).Error
	
	if err != nil {
		return rekening, dto.ErrResponse{Message: "Failed to Create History"}
	}

	return rekening, dto.ErrResponse{Message: ""}
}

func (s *rekeningRepository) TransferRekening(no_rekening uuid.UUID, balance int, pin string, id int) dto.ErrResponse {
	var rekening entity.RekeningEntity
	var user entity.UserEntity
	var rekening_tujuan entity.RekeningEntity
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Where("id = ?", id).First(&rekening).Error
	if rekening.Balance-balance < 0 {
		tx.Rollback()
		return dto.ErrResponse{Message: "Insufficient Balance"}
	}
	if err != nil {
		tx.Rollback()
		return dto.ErrResponse{Message: "There is no rekening with this id"}
	}
	if rekening.Chance > 3 {
		tx.Rollback()
		return dto.ErrResponse{Message: "You have reached the maximum chance"}
	}
	err = bcrypt.CompareHashAndPassword([]byte(rekening.Pin), []byte(pin))
	if err != nil {
		err = tx.Model(&rekening).Update("chance", rekening.Chance+1).Error
		if err != nil {
			tx.Rollback()
			return dto.ErrResponse{Message: "Failed to Update Rekening"}
		}
		tx.Rollback()
		return dto.ErrResponse{Message: "Wrong Pin"}
	}

	err = tx.Model(&rekening).Updates(entity.RekeningEntity{
		Balance: rekening.Balance - balance,
		Chance:  0,
	}).Error
	if err != nil {
		tx.Rollback()
		return dto.ErrResponse{Message: "Failed to Update Rekening"}
	}

	err = tx.Model(&rekening_tujuan).Where("no_rekening = ?", no_rekening).Find(&rekening_tujuan).Error
	if err != nil {
		tx.Rollback()
		return dto.ErrResponse{Message: "There is no rekening with this no_rekening"}
	}

	err = tx.Model(&user).Where("id = ?", rekening_tujuan.UserID).Error
	if err != nil {
		tx.Rollback()
		return dto.ErrResponse{Message: "There is no user with this id"}
	}

	balanceNow := rekening_tujuan.Balance + balance
	err = tx.Model(&entity.RekeningEntity{}).Where("no_rekening =?", no_rekening).Update("balance", balanceNow).Error
	if err != nil {
		tx.Rollback()
		return dto.ErrResponse{Message: "Failed to Update Rekening"}
	}

	err = tx.Model(&entity.HistoryEntity{}).Create(&entity.HistoryEntity{
		Action: "Rekening from " + rekening.NoRekening.String() + " has been transfered to " + rekening_tujuan.NoRekening.String() + " with " + strconv.Itoa(balance) + " Rupiah",
	}).Error

	if err != nil {
		return dto.ErrResponse{Message: "Failed to Create History"}
	}

	err = tx.Commit().Error
	if err != nil {
		return dto.ErrResponse{Message: "Failed because commit error"}
	}

	return dto.ErrResponse{Message: ""}
}
