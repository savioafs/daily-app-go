package repository

import (
	"savioafs/daily-diet-app-go/internal/entity"
)

type UserStorer interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	// Delete User - With Delete all meals of user
	// Get UseProfile (não aqui, mas para puxar as metricas e tudo)
}

type MealStorer interface {
	Create(meal *entity.Meal) (string, error)
	GetMealByID(id string) (*entity.Meal, error)
	GetAllMealsByUser(userID string) ([]entity.Meal, error)
	UpdateMeal(id string, meal *entity.Meal) error

	// DeleteMeal(id string) error
}
