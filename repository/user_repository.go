package repository

import (
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Index() ([]domain.User, error) {
	var users []domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Show(id int) (domain.User, error){
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) Store(user *domain.User) (*domain.User, error) {
	if err := r.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(id int, data *domain.User) (*domain.User,error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	
	if err := r.DB.Model(&user).Updates(data).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Delete(id int) error {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return err
	}
	return r.DB.Delete(&user).Error
}