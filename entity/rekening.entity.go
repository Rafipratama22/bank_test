package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RekeningEntity struct {
	ID         int       `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	NoRekening uuid.UUID `json:"no_rekening" gorm:"type:uuid;"`
	Balance    int       `json:"balance" gorm:"type:int;;default:0"`
	Pin        string    `json:"pin" gorm:"type:varchar(400);"`
	Chance     int       `json:"chance" gorm:"type:int;default:0"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	CreatedAt  time.Time `json:"created_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
}

func (u *RekeningEntity) BeforeCreate(scope *gorm.DB) error {
	u.NoRekening = uuid.New()
	u.CreatedAt = time.Now()
	return nil
}

func (u *RekeningEntity) BeforeUpdate(scope *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
