package model

import "time"

type User struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Password string `db:"password"`
	Meals    []Meal `db:"meals"`
}

type Meal struct {
	ID          string    `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Date        time.Time `db:"date"`
	IsDiet      bool      `db:"is_diet"`
}
