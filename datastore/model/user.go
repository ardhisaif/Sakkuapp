package model

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" gorm:"primary_key"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()

	return
}