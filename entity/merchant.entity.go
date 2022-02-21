package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// swagger:model Merchant
type MerchantEntity struct {
	// Post ID
	// @PrimaryKey
	// @Column(type:uuid, unique: true)
	ID         uuid.UUID `json:"id,omitempty" gorm:"primary_key;auto_incrementype:uuid"`
	// Post Name
	// @Column(type:string)
	Name       string    `json:"name" gorm:"type:varchar(400);"`
	// Post Address
	// @Column(type:string)
	Address    string    `json:"address" gorm:"type:varchar(400);"`
	// Post Phone
	// @Column(type:string)
	Phone      string    `json:"phone" gorm:"type:varchar(400);"`
	// Post No Rekening
	// @Column(type:uuid)
	NoRekening uuid.UUID `json:"no_rekening" gorm:"type:uuid;"`
	// Post User ID
	// @Column(type:uuid)
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;"`
	// Post Created At
	// @Column(type:datetime)
	CreatedAt  time.Time `json:"created_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
	// Post Updated At
	// @Column(type:datetime)
	UpdatedAt  time.Time `json:"updated_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
}

func (u *MerchantEntity) BeforeCreate(scope *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
