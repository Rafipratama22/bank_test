package dto

import (
	"github.com/Rafipratama22/bank_test/entity"
	"github.com/google/uuid"
)


type MerchantListResponse struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	NoRekening uuid.UUID `json:"no_rekening"`
	UserID uuid.UUID `json:"user_id"`
	User entity.UserEntity `json:"user"`
	Rekneing entity.RekeningEntity `json:"rekneing"`
}