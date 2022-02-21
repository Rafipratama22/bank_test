package entity

import (
	"time"

	"gorm.io/gorm"
)

type HistoryEntity struct {
	// Post ID
	// @PrimaryKey
	// @Column(type:int, unique: true, autoincrement: true)
	ID int `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	// Post Action
	// @Column(type:string)
	Action string `json:"action" gorm:"type:varchar(400);not null"`
	// Post Created At
	// @Column(type:datetime)
	CreatedAt time.Time `json:"created_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
	// Post Updated At
	// @Column(type:datetime)
	UpdatedAt time.Time `json:"updated_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
}

func (u *HistoryEntity) BeforeCreate(scope *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
