package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserRepository interface {
	Create(user *User) error
	FindAll() ([]User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}
