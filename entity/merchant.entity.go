package entity

import (
	"time"

	"github.com/google/uuid"
)

type Merchant struct {
	ID         int        `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Name       string     `json:"name" gorm:"type:varchar(400);not null"`
	NoRekening uuid.UUID  `json:"no_rekening" gorm:"type:varchar(100);not null"`
	UserEntity UserEntity `json:"user_entity" gorm:"foreignkey:UserID"`
	UserID     uuid.UUID  `json:"user_id" gorm:"type:uuid;not null;"`
	CreatedAt  time.Time  `json:"created_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
}
