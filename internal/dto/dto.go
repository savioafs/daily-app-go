package dto

import "time"

type MealInputDTO struct {
	Name        string    `json:"name"`
	UserID      string    `json:"user_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsDiet      bool      `json:"is_diet"`
}

type MealOutputDTO struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsDiet      bool      `json:"is_diet"`
}

type UserInputDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOutputDTO struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Email string          `json:"email"`
	Meals []MealOutputDTO `json:"meals"`
}

type MetricsOutputDTO struct {
	TotalMeals        int     `json:"total_meals"`
	TotalMealsDiet    int     `json:"total_meals_diet"`
	TotalMealsNonDiet int     `json:"total_meals_non_diet"`
	DietPercent       float64 `json:"diet_percent"`
	NonDietPercent    float64 `json:"non_diet_percent"`
}
