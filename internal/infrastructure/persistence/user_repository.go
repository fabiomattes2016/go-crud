package persistence

import (
	"github.com/fabiomattes2016/go-crud/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User

	err := r.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User

	err := r.DB.Where("email = ?", email).First(&user).Error

	return &user, err
}

func (r *UserRepository) Update(user *domain.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.User{}, id).Error
}
