package repository

import (
	"database/sql"
	"savioafs/daily-diet-app-go/internal/model"
)

type MealRepositoryPG struct {
	DB *sql.DB
}

func NewMealRepositoryPG(db *sql.DB) *MealRepositoryPG {
	return &MealRepositoryPG{DB: db}
}

func (r *MealRepositoryPG) Create(meal model.Meal) error {
	return nil
}
