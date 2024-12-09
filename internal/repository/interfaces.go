package repository

import (
	"savioafs/daily-diet-app-go/internal/entity"
)

type User interface {
	// Create User
	// UpdateUser
	// Delete User - With Delete all meals of user
	// Get UseProfile (não aqui, mas para puxar as metricas e tudo)
}

type MealStorer interface {
	Create(meal *entity.Meal) (string, error)
	GetMealByID(id string) (*entity.Meal, error)
	GetAllMealsByUser(user_id string) ([]entity.Meal, error)

	// DeleteMeal(id string) error
}
