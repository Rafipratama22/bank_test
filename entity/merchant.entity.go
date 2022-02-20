package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MerchantEntity struct {
	ID         uuid.UUID `json:"id,omitempty" gorm:"primary_key;auto_incrementype:uuid"`
	Name       string    `json:"name" gorm:"type:varchar(400);"`
	Address    string    `json:"address" gorm:"type:varchar(400);"`
	Phone      string    `json:"phone" gorm:"type:varchar(400);"`
	NoRekening uuid.UUID `json:"no_rekening" gorm:"type:uuid;"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
}

func (u *MerchantEntity) BeforeCreate(scope *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
