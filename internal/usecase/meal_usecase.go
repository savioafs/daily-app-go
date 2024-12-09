package usecase

import (
	"errors"
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

func (u *MealUsecase) FindMealByID(id string) (*entity.Meal, error) {
	if id == "" {
		return nil, errors.New("meal id cannot empty")
	}

	meal, err := u.repository.GetMealByID(id)
	if err != nil {
		return nil, err
	}

	if meal == nil {
		return nil, errors.New("meal not found")
	}

	return meal, nil
}

func (u *MealUsecase) GetAllMealsByUser(user_id string) ([]entity.Meal, error) {
	if user_id == "" {
		return nil, errors.New("user id cannot empty")
	}

	meals, err := u.repository.GetAllMealsByUser(user_id)
	if err != nil {
		return nil, err
	}

	if meals == nil {
		return nil, errors.New("meals by id not found")
	}

	return meals, nil
}
