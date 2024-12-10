package usecase

import (
	"savioafs/daily-diet-app-go/internal/entity"
	"savioafs/daily-diet-app-go/internal/repository"
)

type UserUseCase struct {
	repository repository.UserStorer
}

func NewUserUseCase(repo repository.UserStorer) UserUseCase {
	return UserUseCase{repository: repo}
}

func (u *UserUseCase) CreateUser(user *entity.User) error {
	// find user by email para ver se já não tem um email igual
	err := u.repository.Create(user)
	if err != nil {
		return err
	}

	return nil
}
