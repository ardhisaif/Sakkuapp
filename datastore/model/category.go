package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key"`
	UserID      string    `json:"user_id" gorm:"ForeignKey:user_id"`
	Category    string    `json:"category"`
	Type        int8      `json:"type"`
	Total       float64   `json:"total"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (b *Category) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewV4()

	return
}
