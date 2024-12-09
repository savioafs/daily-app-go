package model

import "time"

type Meal struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Date        time.Time `db:"date"`
	IsDiet      bool      `db:"is_diet"`
}
