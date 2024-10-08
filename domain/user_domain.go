package domain

import (
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" validate:"required,min=3,max=50"`
	Email     string    `json:"email" validate:"required,email,min=3,max=100" gorm:"unique"`
	Password  string    `json:"password" validate:"required,min=8,max=255"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	Folders []Folder `gorm:"foreignKey:UserID"`
	Notes []Note `gorm:"foreignKey:UserID"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type UserRepository interface {
	Index() ([]User, error)
	Show(id int) (User, error)
	FindByEmail(email string) (User, error)
	Store(user *User) (*User, error)
	Update(id int, data *User) (*User, error)
	Delete(id int) error
}