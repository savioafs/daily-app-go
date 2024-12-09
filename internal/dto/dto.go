package dto

import "time"

type MealInputDTO struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsDiet      bool      `json:"is_diet"`
}

type MealOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsDiet      bool      `json:"is_diet"`
}
