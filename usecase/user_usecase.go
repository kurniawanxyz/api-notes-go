package usecase

import (
	"github.com/kurniawanxyz/crud-notes-go/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase (repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Index() ([]domain.User, error) {
	return uc.repo.Index()
}

func (uc *UserUseCase) Show(id int) (domain.User, error) {
	return uc.repo.Show(id)
}

func (uc *UserUseCase) Store(user domain.User) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	return uc.repo.Store(&user)
}

func (uc *UserUseCase) Update(id int, data *domain.User) (*domain.User, error) {
	return uc.repo.Update(id, data)
}

func (uc *UserUseCase) Delete(id int) error {
	return uc.repo.Delete(id)
}


