package repository

import "savioafs/daily-diet-app-go/internal/model"

type MealStorer interface {
	Create(meal model.Meal) error
}
