package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PlannedPayment struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	UserID      string    `json:"user_id" gorm:"ForeignKey:user_id"`
	Description string    `json:"description"`
	CategoryID  string    `json:"category_id"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func (b *PlannedPayment) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()

	return
}
