package repository

import (
	"database/sql"
	"savioafs/daily-diet-app-go/internal/entity"

	"github.com/google/uuid"
)

type MealRepositoryPG struct {
	DB *sql.DB
}

func NewMealRepositoryPG(db *sql.DB) *MealRepositoryPG {
	return &MealRepositoryPG{DB: db}
}

func (r *MealRepositoryPG) Create(meal *entity.Meal) (string, error) {
	mealID := uuid.NewString()
	var id string

	stmt, err := r.DB.Prepare("INSERT INTO meals (id, user_id, name, description, date, is_diet) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return "", err
	}

	err = stmt.QueryRow(mealID, meal.UserID, meal.Name, meal.Description, meal.Date, meal.IsDiet).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
