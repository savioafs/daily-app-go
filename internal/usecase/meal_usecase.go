package usecase

import "savioafs/daily-diet-app-go/internal/repository"

type MealUsecase struct {
	repository repository.MealStorer
}

func NewMealUseCase(repo repository.MealStorer) MealUsecase {
	return MealUsecase{repository: repo}
}
