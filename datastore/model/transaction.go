package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	UserID      string    `json:"user_id" gorm:"ForeignKey:user_id"`
	Description string    `json:"description"`
	CategoryID  string    `json:"category_id"`
	Income      float64   `json:"income"`
	Expense     float64   `json:"expense"`
	CreatedAt   time.Time `json:"created_at"`
}

func (b *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()

	return
}
