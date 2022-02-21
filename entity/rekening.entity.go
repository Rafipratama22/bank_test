package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// swagger:model Rekening
type RekeningEntity struct {
	// Post ID
	// @PrimaryKey
	// @Column(type:int, unique: true, autoincrement: true)
	ID         int       `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	// Post No Rekening
	// @Column(type:uuid, unique: true)
	NoRekening uuid.UUID `json:"no_rekening" gorm:"type:uuid;unique"`
	// Post Balance
	// @Column(type:int)
	Balance    int       `json:"balance" gorm:"type:int;;default:0"`
	// Post Pin
	// @Column(type:string)
	Pin        string    `json:"pin" gorm:"type:varchar(400);"`
	// Post Chance
	// @Column(type:datetime)
	Chance     int       `json:"chance" gorm:"type:int;default:0"`
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

func (u *RekeningEntity) BeforeCreate(scope *gorm.DB) error {
	u.NoRekening = uuid.New()
	u.CreatedAt = time.Now()
	return nil
}

func (u *RekeningEntity) BeforeUpdate(scope *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
