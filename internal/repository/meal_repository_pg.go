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

	defer stmt.Close()

	err = stmt.QueryRow(mealID, meal.UserID, meal.Name, meal.Description, meal.Date, meal.IsDiet).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *MealRepositoryPG) GetMealByID(id string) (*entity.Meal, error) {
	stmt, err := r.DB.Prepare("SELECT id, user_id, name, description, date, is_diet FROM meals WHERE id = $1 ")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var meal entity.Meal

	err = stmt.QueryRow(id).Scan(
		&meal.ID,
		&meal.UserID,
		&meal.Name,
		&meal.Description,
		&meal.Date,
		&meal.IsDiet)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &meal, nil
}
