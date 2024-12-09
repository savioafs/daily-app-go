package usecase

import (
	"savioafs/daily-diet-app-go/internal/entity"
	"savioafs/daily-diet-app-go/internal/repository"
)

type MealUsecase struct {
	repository repository.MealStorer
}

func NewMealUseCase(repo repository.MealStorer) MealUsecase {
	return MealUsecase{repository: repo}
}

func (u *MealUsecase) Create(meal *entity.Meal) (*entity.Meal, error) {
	mealID, err := u.repository.Create(meal)
	if err != nil {
		return nil, err
	}

	meal.ID = mealID

	return meal, nil
}
