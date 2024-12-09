package repository

import (
	"savioafs/daily-diet-app-go/internal/entity"
)

type MealStorer interface {
	Create(meal *entity.Meal) (string, error)
}
