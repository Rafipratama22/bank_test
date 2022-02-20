package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	Name      string    `json:"name" gorm:"type:varchar(400);not null"`
	Password  string    `json:"password" gorm:"type:varchar(400);not null"`
	Email     string    `json:"email" gorm:"type:varchar(400);not null"`
	IsActive  bool      `json:"is_active" gorm:"type:boolean;not null;default:true"`
	// Merchant  *MerchantEntity
	CreatedAt time.Time `json:"created_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
}

func (u *UserEntity) BeforeCreate(scope *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	return nil
}

func (u *UserEntity) BeforeUpdate(scope *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
