package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Income      float64   `json:"income"`
	Expense     float64   `json:"expense"`
	CreatedAt   time.Time `json:"created_at"`
}

func (b *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()

	return
}
