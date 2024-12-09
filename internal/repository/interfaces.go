package repository

import (
	"savioafs/daily-diet-app-go/internal/entity"
)

type MealStorer interface {
	Create(meal *entity.Meal) (string, error)
	GetMealByID(id string) (*entity.Meal, error)
	// GetAllMealsByUser(userID string) ([]entity.Meal, error)
	// DeleteMeal(id string) error
}
