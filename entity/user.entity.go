package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserEntity struct {
	// Post ID
	// @PrimaryKey
	// @Column(type:uuid, unique: true)
	ID        uuid.UUID `json:"id,omitempty" gorm:"type:uuid;primary_key"`
	// Post Name
	// @Column(type:string)
	Name      string    `json:"name" gorm:"type:varchar(400);not null"`
	// Post Email
	// @Column(type:string)
	Password  string    `json:"password" gorm:"type:varchar(400);not null"`
	// Post Email
	// @Column(type:string)
	Email     string    `json:"email" gorm:"type:varchar(400);not null"`
	// Post IsActive
	// @Column(type:boolean)
	IsActive  bool      `json:"is_active" gorm:"type:boolean;not null;default:true"`
	// Post Created At
	// @Column(type:datetime)
	CreatedAt time.Time `json:"created_at" gorm:"type:time;default:CURRENT_TIMESTAMP"`
	// Post Updated At
	// @Column(type:datetime)
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
