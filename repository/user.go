package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"primaryKey,not null"`
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"not null,unique"`
	Password string    `json:"password" gorm:"not null"`
	Role     string    `json:"role"`
}

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	Insert(params *User) (*User, error)
}
