package entity

import "time"

type HistoryEntity struct {
	ID int `json:"id,omitempty" gorm:"primary_key;auto_increment"`
	Action string `json:"action" gorm:"type:varchar(400);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:date;default:CURRENT_TIMESTAMP"`
}