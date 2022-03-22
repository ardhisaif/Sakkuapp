package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Balance struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key"`
	UserID    uuid.UUID `json:"user_id" gorm:"ForeignKey:user_id"`
	Balance   float64   `json:"balance"`
	Expense   float64   `json:"expense"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *Balance) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()

	return
}
